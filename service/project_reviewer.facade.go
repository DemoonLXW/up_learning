package service

import (
	"encoding/json"
	"fmt"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/entity/workflow"
)

type ProjectReviewerFacade struct {
	Workflow  *WorkflowService
	Applicant *ApplicantService
	DB        *ent.Client
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

func (faca *ProjectReviewerFacade) FindReviewProjectTaskDetailByID(id string) (*workflow.HistoricTaskInstances, error) {
	body := map[string]interface{}{
		"processDefinitionKey":    workflow.DefinitionReviewProject,
		"includeProcessVariables": true,
		"taskId":                  id,
	}

	b, err := faca.Workflow.QueryForHistoricTaskInstances(body)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project detail by id failed: %w", err)
	}
	taskList := workflow.HistoricTaskInstancesPageList{}
	err = json.Unmarshal(b, &taskList)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project detail by id failed: %w", err)
	}
	if len(taskList.Data) == 0 {
		return nil, fmt.Errorf("retrieve review project detail by id not found")
	}

	return &taskList.Data[0], nil
}

func (faca *ProjectReviewerFacade) ReviewProjectBytaskID(taskID string, reviewerID uint32, action uint8) error {
	body := map[string]interface{}{
		"action": "complete",
		"variables": []interface{}{
			map[string]interface{}{
				"name":  "reviewer",
				"value": reviewerID,
			},
			map[string]interface{}{
				"name":  "action",
				"value": action,
			},
		},
	}

	err := faca.Workflow.TaskActions(taskID, body)
	if err != nil {
		return fmt.Errorf("retrieve review project detail by id failed: %w", err)
	}

	return nil
}

func (faca *ProjectReviewerFacade) RetrieveReviewProjectTaskHistoryByUserID(userID uint32, search *entity.SearchReviewProjectTask) (*workflow.HistoricTaskInstancesPageList, error) {
	body := map[string]interface{}{
		"processDefinitionKey":      workflow.DefinitionReviewProject,
		"includeProcessVariables":   true,
		"includeTaskLocalVariables": true,
		"taskVariables": []interface{}{
			map[string]interface{}{
				"name":      "reviewer",
				"value":     userID,
				"operation": "equals",
				"type":      "integer",
			},
		},
	}

	if search.PageSize != nil && search.Current != nil {
		offset := int((*search.Current - 1) * (*search.PageSize))
		body["start"] = offset
		body["size"] = *search.PageSize
	}
	isSorted := search.Sort != "" && (search.Sort == "startTime")
	if isSorted && search.Order != nil {
		if *search.Order {
			body["sort"] = search.Sort
			body["order"] = "desc"
		} else {
			body["sort"] = search.Sort
			body["order"] = "asc"
		}
	}

	b, err := faca.Workflow.QueryForHistoricTaskInstances(body)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project task history failed: %w", err)
	}

	list := workflow.HistoricTaskInstancesPageList{}
	err = json.Unmarshal(b, &list)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project task history failed: %w", err)
	}

	return &list, nil
}
