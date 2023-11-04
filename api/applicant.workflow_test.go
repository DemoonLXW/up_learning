package api_test

import (
	"os"
	"testing"
	"time"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/workflow"
	"github.com/stretchr/testify/assert"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/testsuite"
	"go.uber.org/zap"
)

func CreateWorkflowEnvironment() (*testsuite.TestWorkflowEnvironment, error) {
	os.Chdir("../")
	os.Setenv("DB_CONFIG", "./database.config.json")
	os.Setenv("DOMAIN_CONFIG", "./domain.config.json")
	// os.Setenv("GIN_LOG", "./gin.log")
	// os.Setenv("WORKFLOW_CONFIG", "./workflow.config.yaml")

	testSuite := new(testsuite.WorkflowTestSuite)
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	testSuite.SetLogger(logger)
	env := testSuite.NewTestWorkflowEnvironment()
	return env, nil
}

func TestReviewProjectWorkflow(t *testing.T) {
	env, err := CreateWorkflowEnvironment()
	assert.Nil(t, err)

	cont, err := injection.ProvideController()
	assert.Nil(t, err)

	toAdd := entity.ToAddReviewProject{
		ApplicantID: uint32(32),
		ProjectID:   uint32(322),
		WorkflowID:  "work flow id",
		RunID:       "run id",
	}

	env.ExecuteWorkflow(cont.Applicant.ReviewProjectWorkflow, &toAdd)

	assert.True(t, env.IsWorkflowCompleted())
	assert.NoError(t, env.GetWorkflowError())
}
func TestWorkflow(t *testing.T) {
	os.Chdir("../")

	cont, err := injection.ProvideController()
	assert.Nil(t, err)

	helper, err := injection.ProvideWorkflowHelper()
	assert.Nil(t, err)

	toAdd := entity.ToAddReviewProject{
		ApplicantID: uint32(32),
		ProjectID:   uint32(322),
	}

	helper.RegisterWorkflow(cont.Applicant.ReviewProjectWorkflow)
	helper.RegisterActivity(cont.Applicant.Applicant.FindUserWithTeacherOrStudentById)

	_, err = helper.StartWorkflow(
		client.StartWorkflowOptions{
			ID:                           "start1",
			TaskList:                     workflow.ReviewProjectTask,
			ExecutionStartToCloseTimeout: 120 * time.Second,
		},
		cont.Applicant.ReviewProjectWorkflow,
		toAdd.ProjectID,
		toAdd.ApplicantID,
	)
	assert.Nil(t, err)

}
