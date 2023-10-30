package workflow

import (
	"github.com/google/wire"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

var WorkflowHelperProvider = wire.NewSet(
	ProvideWorkflowHelper,
)

const (
	ReviewProjectTask = "review-project-list"
)

func SetupWorkflow(wh *WorkflowHelper) error {

	tasks := []string{ReviewProjectTask}
	errChan := make(chan error)

	for i := range tasks {
		cfg := zap.NewProductionConfig()
		cfg.OutputPaths = []string{tasks[i] + ".log"}
		logger, err := cfg.Build()
		if err != nil {
			return err
		}

		go func(c chan<- error, l *zap.Logger) {
			workerOptions := worker.Options{
				Logger: l,
			}
			err = wh.StartWorkers(wh.Config.DomainName, tasks[i], workerOptions)
			errChan <- err
			if err == nil {
				select {}
			}
		}(errChan, logger)
		err = <-errChan
		if err != nil {
			return err
		}
	}

	return nil
}
