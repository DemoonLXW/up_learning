package service

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/database/ent/projectfile"
	"github.com/DemoonLXW/up_learning/entity"
)

type ApplicantFacade struct {
	Common   *CommonService
	Workflow *WorkflowService
	DB       *ent.Client
}

func (faca *ApplicantFacade) createProjectAttachments(ctx context.Context, client *ent.Client, attachments []*entity.ToAddFile, projectID uint32) error {
	fileIDs, err := faca.Common.CreateFile(ctx, client, attachments)
	if err != nil {
		return fmt.Errorf("create project attachments create file failed: %w", err)
	}

	fileNum := len(fileIDs)
	bulk := make([]*ent.ProjectFileCreate, fileNum)
	for i := 0; i < fileNum; i++ {
		bulk[i] = client.ProjectFile.Create().SetFid(fileIDs[i]).SetPid(projectID)
	}

	creates, err := client.ProjectFile.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return fmt.Errorf("create project attachments add attachments: %w", err)
	}
	if len(creates) != fileNum {
		return fmt.Errorf("create project attachments add attachments: some files can not be found")
	}
	return nil
}

func (faca *ApplicantFacade) CreateProject(ctx context.Context, client *ent.Client, toCreates []*entity.ToAddProject) error {
	if ctx == nil || client == nil {
		return fmt.Errorf("context or client is nil")
	}

	current := 0
	length := len(toCreates)
	for ; current < length; current++ {
		id, err := client.Project.Query().Where(func(s *sql.Selector) {
			s.Where(sql.NotNull(project.FieldDeletedTime))
		}).FirstID(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return fmt.Errorf("create project find a deleted project id query failed: %w", err)
			}
			break
		}

		uNum, err := client.Project.Update().Where(project.And(
			project.IDEQ(id),
			func(s *sql.Selector) {
				s.Where(sql.NotNull(project.FieldDeletedTime))
			},
		)).
			SetCreatedTime(time.Now()).
			ClearModifiedTime().
			ClearDeletedTime().
			SetUID(toCreates[current].UID).
			SetTitle(toCreates[current].Title).
			SetGoal(toCreates[current].Goal).
			SetPrinciple(toCreates[current].Principle).
			SetProcessAndMethod(toCreates[current].ProcessAndMethod).
			SetStep(toCreates[current].Step).
			SetResultAndConclusion(toCreates[current].ResultAndConclusion).
			SetRequirement(toCreates[current].Requirement).
			SetReviewStatus(uint8(0)).
			SetIsDisabled(false).
			Save(ctx)

		if err != nil {
			return fmt.Errorf("create project: %w", err)
		}
		if uNum == 0 {
			return fmt.Errorf("create project update deleted project affect 0 row")
		}

		if toCreates[current].Attachments != nil {
			// fileIDs, err := faca.Common.CreateFile(ctx, client, toCreates[current].Attachments)
			// if err != nil {
			// 	return fmt.Errorf("create project create file failed: %w", err)
			// }

			// fileNum := len(fileIDs)
			// bulk := make([]*ent.ProjectFileCreate, fileNum)
			// for i := 0; i < fileNum; i++ {
			// 	bulk[i] = client.ProjectFile.Create().SetFid(fileIDs[i]).SetPid(id)
			// }

			// creates, err := client.ProjectFile.CreateBulk(bulk...).Save(ctx)
			// if err != nil {
			// 	return fmt.Errorf("create project add attachments: %w", err)
			// }
			// if len(creates) != fileNum {
			// 	return fmt.Errorf("create project add attachments: some files can not be found")
			// }
			err := faca.createProjectAttachments(ctx, client, toCreates[current].Attachments, id)
			if err != nil {
				return fmt.Errorf("create project attachments failed: %w", err)
			}
		}
	}
	if current < length {
		num, err := client.Project.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create project count failed: %w", err)
		}

		bulkLength := length - current
		// bulk := make([]*ent.ProjectCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			projectID := uint32(num + i + 1)

			err := client.Project.Create().
				SetID(projectID).
				SetUID(toCreates[current+i].UID).
				SetTitle(toCreates[current+i].Title).
				SetGoal(toCreates[current+i].Goal).
				SetPrinciple(toCreates[current+i].Principle).
				SetProcessAndMethod(toCreates[current+i].ProcessAndMethod).
				SetStep(toCreates[current+i].Step).
				SetResultAndConclusion(toCreates[current+i].ResultAndConclusion).
				SetRequirement(toCreates[current+i].Requirement).
				Exec(ctx)

			if err != nil {
				return fmt.Errorf("create project failed: %w", err)
			}

			if toCreates[current+i].Attachments != nil {
				// fileIDs, err := faca.Common.CreateFile(ctx, client, toCreates[current+i].Attachments)
				// if err != nil {
				// 	return fmt.Errorf("create project create file failed: %w", err)
				// }

				// fileNum := len(fileIDs)
				// bulk := make([]*ent.ProjectFileCreate, fileNum)
				// for i := 0; i < fileNum; i++ {
				// 	bulk[i] = client.ProjectFile.Create().SetFid(fileIDs[i]).SetPid(projectID)
				// }
				// creates, err := client.ProjectFile.CreateBulk(bulk...).Save(ctx)

				// if err != nil {
				// 	return fmt.Errorf("create project add attachments: %w", err)
				// }
				// if len(creates) != fileNum {
				// 	return fmt.Errorf("create project add attachments: some files can not be found")
				// }
				err := faca.createProjectAttachments(ctx, client, toCreates[current+i].Attachments, projectID)
				if err != nil {
					return fmt.Errorf("create project attachments failed: %w", err)
				}
			}
		}
	}

	return nil
}

func (faca *ApplicantFacade) DeleteProject(ctx context.Context, client *ent.Client, toDeletes []*ent.Project) error {
	if ctx == nil || client == nil {
		return fmt.Errorf("context or client is nil")
	}

	for i := range toDeletes {
		attachmentIDs := make([]uint32, 0)
		for _, a := range toDeletes[i].Edges.Attachments {
			attachmentIDs = append(attachmentIDs, a.ID)
		}
		toDeleteFiles, err := client.File.Query().Where(file.And(
			file.IDIn(attachmentIDs...),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(file.FieldDeletedTime))
			},
		)).All(ctx)
		if err != nil || len(toDeleteFiles) != len(attachmentIDs) {
			return fmt.Errorf("delete project delete attachments not found files failed: %w", err)
		}

		err = faca.Common.DeleteFile(ctx, client, toDeleteFiles)
		if err != nil {
			return fmt.Errorf("delete project delete attachments failed: %w", err)
		}

		num, err := client.Project.Update().
			Where(project.And(
				project.IDEQ(toDeletes[i].ID),
				func(s *sql.Selector) {
					s.Where(sql.IsNull(project.FieldDeletedTime))
				},
			)).
			ClearModifiedTime().
			ClearAttachments().
			SetDeletedTime(time.Now()).
			SetTitle("*" + toDeletes[i].Title).
			Save(ctx)

		if err != nil {
			return fmt.Errorf("delete project failed: %w", err)
		}
		if num == 0 {
			return fmt.Errorf("delete project affect 0 row: id[%d]", toDeletes[i].ID)
		}

	}

	return nil
}

func (faca *ApplicantFacade) UpdateProject(ctx context.Context, client *ent.Client, toUpdate *entity.ToModifyProject) error {
	if ctx == nil || client == nil {
		return fmt.Errorf("context or client is nil")
	}

	updater := client.Project.Update().Where(
		project.IDEQ(toUpdate.ID),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(project.FieldDeletedTime))
		},
	)
	if toUpdate.Title != nil {
		updater.SetTitle(*toUpdate.Title)
	}
	if toUpdate.Goal != nil {
		updater.SetGoal(*toUpdate.Goal)
	}
	if toUpdate.Principle != nil {
		updater.SetPrinciple(*toUpdate.Principle)
	}
	if toUpdate.ProcessAndMethod != nil {
		updater.SetProcessAndMethod(*toUpdate.ProcessAndMethod)
	}
	if toUpdate.Step != nil {
		updater.SetStep(*toUpdate.Step)
	}
	if toUpdate.ResultAndConclusion != nil {
		updater.SetResultAndConclusion(*toUpdate.ResultAndConclusion)
	}
	if toUpdate.Requirement != nil {
		updater.SetRequirement(*toUpdate.Requirement)
	}
	if toUpdate.ReviewStatus != nil {
		updater.SetReviewStatus(*toUpdate.ReviewStatus)
	}
	updater.SetModifiedTime(time.Now())

	uNum, err := updater.Save(ctx)
	if err != nil {
		return fmt.Errorf("update project failed: %w", err)
	}
	if uNum == 0 {
		return fmt.Errorf("update project affect 0 row")
	}

	if toUpdate.DeleteFileIDs != nil {
		// Ensure that the deleted attachment belongs to the project
		dNum, err := client.ProjectFile.Delete().Where(
			projectfile.FidIn(toUpdate.DeleteFileIDs...),
			projectfile.Pid(toUpdate.ID),
		).Exec(ctx)
		if err != nil {
			return fmt.Errorf("update project delete attachments relationship: %w", err)
		}
		if dNum != len(toUpdate.DeleteFileIDs) {
			return fmt.Errorf("update project delete attachments relationship: some files can not be found")
		}
		// the order is important
		toDeleteFiles, err := client.File.Query().Where(file.And(
			file.IDIn(toUpdate.DeleteFileIDs...),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(file.FieldDeletedTime))
			},
		)).All(ctx)
		if err != nil || len(toDeleteFiles) != len(toUpdate.DeleteFileIDs) {
			return fmt.Errorf("update project delete attachments not found files failed: %w", err)
		}
		err = faca.Common.DeleteFile(ctx, client, toDeleteFiles)
		if err != nil {
			return fmt.Errorf("update project delete attachments failed: %w", err)
		}

	}

	if toUpdate.AddFile != nil {
		// fileIDs, err := faca.Common.CreateFile(ctx, client, toUpdate.AddFile)
		// if err != nil {
		// 	return fmt.Errorf("update project add file failed: %w", err)
		// }

		// fileNum := len(fileIDs)
		// bulk := make([]*ent.ProjectFileCreate, fileNum)
		// for i := 0; i < fileNum; i++ {
		// 	bulk[i] = client.ProjectFile.Create().SetFid(fileIDs[i]).SetPid(toUpdate.ID)
		// }

		// creates, err := client.ProjectFile.CreateBulk(bulk...).Save(ctx)

		// if err != nil {
		// 	return fmt.Errorf("update project add attachments: %w", err)
		// }
		// if len(creates) != fileNum {
		// 	return fmt.Errorf("update project add attachments: some files can not be found")
		// }
		err := faca.createProjectAttachments(ctx, client, toUpdate.AddFile, toUpdate.ID)
		if err != nil {
			return fmt.Errorf("update project attachments failed: %w", err)
		}
	}

	return nil
}

func (faca *ApplicantFacade) StartReviewProjectProcess(userID, projectID uint32) error {

	body := map[string]interface{}{
		"processDefinitionKey": entity.REVIEW_PROJECT,
		"variables": []interface{}{
			map[string]interface{}{
				"name":  "userID",
				"value": userID,
			},
			map[string]interface{}{
				"name":  "projectID",
				"value": projectID,
			},
			map[string]interface{}{
				"name":  "dueDate",
				"value": time.Now().Add(7 * 24 * time.Hour),
			},
		},
	}

	err := faca.Workflow.StartAProcessInstance(body)
	if err != nil {
		return fmt.Errorf("submit project for review failed: %w", err)
	}

	return nil
}

func (faca *ApplicantFacade) RetrieveReviewProjectRecordByProjectID(search *entity.SearchReviewProjectRecord) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"processDefinitionKey":    entity.REVIEW_PROJECT,
		"includeProcessVariables": true,
	}

	variables := []interface{}{
		map[string]interface{}{
			"name":      "projectID",
			"value":     *search.ProjectID,
			"operation": "equals",
			"type":      "integer",
		},
	}
	if search.ReviewStatus != nil {
		variables = append(variables, map[string]interface{}{
			"name":      "reviewStatus",
			"value":     *search.ReviewStatus,
			"operation": "equals",
			"type":      "integer",
		})
	}
	body["variables"] = variables

	// if search.IsFinished != nil {
	// 	body["finished"] = search.IsFinished
	// }
	if search.PageSize != nil && search.Current != nil {
		offset := int((*search.Current - 1) * (*search.PageSize))
		body["start"] = offset
		body["size"] = *search.PageSize
	}

	isSorted := search.Sort != "" && (search.Sort == "processInstanceId" || search.Sort == "startTime" || search.Sort == "endTime")
	if isSorted && search.Order != nil {
		if *search.Order {
			body["sort"] = search.Sort
			body["order"] = "desc"
		} else {
			body["sort"] = search.Sort
			body["order"] = "asc"
		}
	}

	m, err := faca.Workflow.QueryForHistoricProcessInstances(body)
	if err != nil {
		return nil, fmt.Errorf("retrieve review project record failed: %w", err)
	}

	return m, nil
}

func (faca *ApplicantFacade) FindReviewProjectRecordDetailById(id string) (map[string]interface{}, map[string]interface{}, error) {
	processbody := map[string]interface{}{
		"processDefinitionKey":    entity.REVIEW_PROJECT,
		"includeProcessVariables": true,
		"processInstanceId":       id,
	}

	processMap, err := faca.Workflow.QueryForHistoricProcessInstances(processbody)
	if err != nil {
		return nil, nil, fmt.Errorf("retrieve review project record detail failed: %w", err)
	}

	taskbody := map[string]interface{}{
		"processDefinitionKey":      entity.REVIEW_PROJECT,
		"includeTaskLocalVariables": true,
		"processInstanceId":         id,
	}

	taskMap, err := faca.Workflow.QueryForHistoricTaskInstances(taskbody)
	if err != nil {
		return nil, nil, fmt.Errorf("retrieve review project record detail failed: %w", err)
	}

	return processMap, taskMap, nil
}
