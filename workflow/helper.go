package workflow

import (
	"context"
	"fmt"
	"os"

	apiv1 "github.com/uber/cadence-idl/go/proto/api/v1"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/compatibility"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/grpc"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

const (
	_cadenceClientName      = "cadence-client"
	_cadenceFrontendService = "cadence-frontend"
)

type (
	WorkflowHelper struct {
		Service       workflowserviceclient.Interface
		CadenceClient client.Client
		Logger        *zap.Logger
		Config        Configuration
		HostPort      string
		Dispatcher    *yarpc.Dispatcher
		Domain        string

		workflowRegistries []registryOption
		activityRegistries []registryOption

		configFile string
	}

	Configuration struct {
		DomainName      string `yaml:"domain"`
		ServiceName     string `yaml:"service"`
		HostNameAndPort string `yaml:"host"`
	}

	registryOption struct {
		registry interface{}
		alias    string
	}
)

func ProvideWorkflowHelper() (*WorkflowHelper, error) {
	var h WorkflowHelper

	defaultConfigFile, ok := os.LookupEnv("WORKFLOW_CONFIG")
	if !ok {
		defaultConfigFile = "./workflow.config.yaml"
	}

	h.configFile = defaultConfigFile

	configData, err := os.ReadFile(defaultConfigFile)
	if err != nil {
		return nil, fmt.Errorf("failed to log config file: %v, Error: %w", defaultConfigFile, err)
	}

	if err := yaml.Unmarshal(configData, &h.Config); err != nil {
		return nil, fmt.Errorf("error initializing configuration: %w", err)
	}

	// Initialize logger for running samples
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	logger.Info("Logger created.")
	h.Logger = logger

	h.HostPort = h.Config.HostNameAndPort
	h.Domain = h.Config.DomainName
	dispatcher, err := h.BuildDispatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to build dispatcher: %w", err)

	}
	h.Dispatcher = dispatcher
	if err := h.Dispatcher.Start(); err != nil {
		return nil, fmt.Errorf("failed to create outbound transport channel: %w", err)
	}

	service, err := h.BuildServiceClient()
	if err != nil {
		return nil, err
	}
	h.Service = service

	cadenceClient, err := h.BuildCadenceClient()
	if err != nil {
		return nil, err
	}
	h.CadenceClient = cadenceClient

	h.workflowRegistries = make([]registryOption, 0, 1)
	h.activityRegistries = make([]registryOption, 0, 1)

	return &h, nil
}

func (h *WorkflowHelper) BuildDispatcher() (*yarpc.Dispatcher, error) {
	if len(h.HostPort) == 0 {
		return nil, fmt.Errorf("host name is empty")
	}

	h.Logger.Debug("Creating RPC Dispatcher outbound",
		zap.String("ServiceName", _cadenceFrontendService),
		zap.String("HostPort", h.HostPort))

	return yarpc.NewDispatcher(yarpc.Config{
		Name: _cadenceClientName,
		Outbounds: yarpc.Outbounds{
			_cadenceFrontendService: {Unary: grpc.NewTransport().NewSingleOutbound(h.HostPort)},
		},
	}), nil
}

func (h *WorkflowHelper) BuildServiceClient() (workflowserviceclient.Interface, error) {

	if h.Dispatcher == nil {
		return nil, fmt.Errorf("no RPC Dispatcher provided to create a connection to Cadence Service")
	}

	clientConfig := h.Dispatcher.ClientConfig(_cadenceFrontendService)
	return compatibility.NewThrift2ProtoAdapter(
		apiv1.NewDomainAPIYARPCClient(clientConfig),
		apiv1.NewWorkflowAPIYARPCClient(clientConfig),
		apiv1.NewWorkerAPIYARPCClient(clientConfig),
		apiv1.NewVisibilityAPIYARPCClient(clientConfig),
	), nil
}

func (h *WorkflowHelper) BuildCadenceClient() (client.Client, error) {

	return client.NewClient(
		h.Service,
		h.Domain,
		&client.Options{},
	), nil
}

func (h *WorkflowHelper) RegisterWorkflow(workflow interface{}) {
	h.RegisterWorkflowWithAlias(workflow, "")
}

func (h *WorkflowHelper) RegisterWorkflowWithAlias(workflow interface{}, alias string) {
	registryOption := registryOption{
		registry: workflow,
		alias:    alias,
	}
	h.workflowRegistries = append(h.workflowRegistries, registryOption)
}

func (h *WorkflowHelper) RegisterActivity(activity interface{}) {
	h.RegisterActivityWithAlias(activity, "")
}

func (h *WorkflowHelper) RegisterActivityWithAlias(activity interface{}, alias string) {
	registryOption := registryOption{
		registry: activity,
		alias:    alias,
	}
	h.activityRegistries = append(h.activityRegistries, registryOption)
}

func (h *WorkflowHelper) registerWorkflowAndActivity(worker worker.Worker) {
	for _, w := range h.workflowRegistries {
		if len(w.alias) == 0 {
			worker.RegisterWorkflow(w.registry)
		} else {
			worker.RegisterWorkflowWithOptions(w.registry, workflow.RegisterOptions{Name: w.alias})
		}
	}
	for _, act := range h.activityRegistries {
		if len(act.alias) == 0 {
			worker.RegisterActivity(act.registry)
		} else {
			worker.RegisterActivityWithOptions(act.registry, activity.RegisterOptions{Name: act.alias})
		}
	}
}

func (h *WorkflowHelper) StartWorkers(domainName string, groupName string, options worker.Options) error {
	worker := worker.New(h.Service, domainName, groupName, options)
	h.registerWorkflowAndActivity(worker)

	err := worker.Start()
	if err != nil {
		// h.Logger.Error("Failed to start workers.", zap.Error(err))
		// panic("Failed to start workers")
		return fmt.Errorf("failed to start workers")
	}
	return nil
}

func (h *WorkflowHelper) StartWorkflow(
	options client.StartWorkflowOptions,
	workflow interface{},
	args ...interface{},
) (*workflow.Execution, error) {
	return h.StartWorkflowWithCtx(context.Background(), options, workflow, args...)
}

// StartWorkflowWithCtx starts a workflow with the provided context
func (h *WorkflowHelper) StartWorkflowWithCtx(
	ctx context.Context,
	options client.StartWorkflowOptions,
	workflow interface{},
	args ...interface{},
) (*workflow.Execution, error) {

	we, err := h.CadenceClient.StartWorkflow(ctx, options, workflow, args...)
	if err != nil {
		// h.Logger.Error("Failed to create workflow", zap.Error(err))
		return nil, fmt.Errorf("failed to create workflow: %w", err)
	} else {
		// h.Logger.Info("Started Workflow", zap.String("WorkflowID", we.ID), zap.String("RunID", we.RunID))
		return we, nil
	}
}

func (h *WorkflowHelper) SignalWorkflow(workflowID, signal string, data interface{}) error {

	err := h.CadenceClient.SignalWorkflow(context.Background(), workflowID, "", signal, data)
	if err != nil {
		// h.Logger.Error("Failed to signal workflow", zap.Error(err))
		return fmt.Errorf("failed to signal workflow: %w", err)
	}

	return nil
}
