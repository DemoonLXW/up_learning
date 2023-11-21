package service

import (
	"encoding/json"
	"fmt"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/entity/workflow"
)

type ProjectReviewerFacade struct {
	Workflow *WorkflowService
	DB       *ent.Client
}

func (faca *ProjectReviewerFacade) RetrieveReviewProjectTaskOfTeacher(search *entity.SearchReviewProjectTask, userID uint32) (*workflow.TaskPageList, error) {
	uid := fmt.Sprintf("%d", userID)
	body := map[string]interface{}{
		"processDefinitionKey":    workflow.DefinitionReviewProject,
		"includeProcessVariables": true,
		"assignee":                uid,
	}
	if search.PageSize != nil && search.Current != nil {
		offset := int((*search.Current - 1) * (*search.PageSize))
		body["start"] = offset
		body["size"] = *search.PageSize
	}
	isSorted := search.Sort != "" && (search.Sort == "id" || search.Sort == "createTime")
	if isSorted && search.Order != nil {
		if *search.Order {
			body["sort"] = search.Sort
			body["order"] = "desc"
		} else {
			body["sort"] = search.Sort
			body["order"] = "asc"
		}
	}

	b, err := faca.Workflow.QueryForTasks(body)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project of teacher failed: %w", err)
	}
	list := workflow.TaskPageList{}
	err = json.Unmarshal(b, &list)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project of teacher failed: %w", err)
	}

	return &list, nil
}

func (faca *ProjectReviewerFacade) RetrieveReviewProjectTaskOfTeacherReviewer(search *entity.SearchReviewProjectTask, collegeID uint8) (*workflow.TaskPageList, error) {
	cid := fmt.Sprintf("%d", collegeID)
	body := map[string]interface{}{
		"processDefinitionKey":    workflow.DefinitionReviewProject,
		"includeProcessVariables": true,
		"candidateGroup":          "teacher project reviewer" + cid,
	}
	if search.PageSize != nil && search.Current != nil {
		offset := int((*search.Current - 1) * (*search.PageSize))
		body["start"] = offset
		body["size"] = *search.PageSize
	}
	isSorted := search.Sort != "" && (search.Sort == "id" || search.Sort == "createTime")
	if isSorted && search.Order != nil {
		if *search.Order {
			body["sort"] = search.Sort
			body["order"] = "desc"
		} else {
			body["sort"] = search.Sort
			body["order"] = "asc"
		}
	}

	b, err := faca.Workflow.QueryForTasks(body)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project of teacher reviewer failed: %w", err)
	}
	list := workflow.TaskPageList{}
	err = json.Unmarshal(b, &list)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project of teacher reviewer failed: %w", err)
	}

	return &list, nil
}

func (faca *ProjectReviewerFacade) RetrieveReviewProjectTaskOfPlatformReviewer(search *entity.SearchReviewProjectTask) (*workflow.TaskPageList, error) {
	body := map[string]interface{}{
		"processDefinitionKey":    workflow.DefinitionReviewProject,
		"includeProcessVariables": true,
		"candidateGroup":          "platform project reviewer",
	}
	if search.PageSize != nil && search.Current != nil {
		offset := int((*search.Current - 1) * (*search.PageSize))
		body["start"] = offset
		body["size"] = *search.PageSize
	}
	isSorted := search.Sort != "" && (search.Sort == "id" || search.Sort == "createTime")
	if isSorted && search.Order != nil {
		if *search.Order {
			body["sort"] = search.Sort
			body["order"] = "desc"
		} else {
			body["sort"] = search.Sort
			body["order"] = "asc"
		}
	}

	b, err := faca.Workflow.QueryForTasks(body)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project of platform reviewer failed: %w", err)
	}
	list := workflow.TaskPageList{}
	err = json.Unmarshal(b, &list)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project of platform reviewer failed: %w", err)
	}

	return &list, nil
}
