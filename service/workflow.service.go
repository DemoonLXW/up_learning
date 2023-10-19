package service

import (
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/workflow"
)

type WorkflowService struct {
	DB *ent.Client
	WH *workflow.WorkflowHelper
}
