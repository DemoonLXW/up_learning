package workflow

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvideWorkflowHelper(t *testing.T) {
	os.Setenv("WORKFLOW_CONFIG", "../workflow.config.yaml")
	_, err := ProvideWorkflowHelper()
	assert.Nil(t, err)
}
