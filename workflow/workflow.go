package workflow

import (
	"github.com/google/wire"
	"go.uber.org/cadence/worker"
)

var WorkflowHelperProvider = wire.NewSet(
	ProvideWorkflowHelper,
)

const (
	ReviewProjectTask = "review-project-list"
)

func SetupWorkflow(wh *WorkflowHelper) error {
	workerOptions := worker.Options{
		Logger: wh.Logger,
	}

	tasks := []string{ReviewProjectTask}
	errChan := make(chan error)

	for i := range tasks {
		go func(c chan<- error) {
			err := wh.StartWorkers(wh.Config.DomainName, tasks[i], workerOptions)
			errChan <- err
			if err == nil {
				select {}
			}
		}(errChan)
		err := <-errChan
		if err != nil {
			return err
		}
	}

	return nil
}
