package service

import (
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/entity"
)

type ProjectReviewerFacade struct {
	Workflow *WorkflowService
	DB       *ent.Client
}

func (faca *ProjectReviewerFacade) RetrieveReviewProjectTaskByTeacher(search *entity.SearchReviewProjectTask) (map[string]interface{}, error) {
	return nil, nil
}
