package facade

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
)

type TeacherFacade struct {
	Common *service.CommonService
}

func (serv *TeacherFacade) CreateProject(ctx context.Context, client *ent.Client, toCreates []*entity.ToAddProject) error {
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

		num, err := client.Project.Update().Where(project.And(
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
		if num == 0 {
			return fmt.Errorf("create project update deleted project affect 0 row")
		}
	}
	if current < length {
		num, err := client.Project.Query().Aggregate(ent.Count()).Int(ctx)
		if err != nil {
			return fmt.Errorf("create project count failed: %w", err)
		}

		bulkLength := length - current
		bulk := make([]*ent.ProjectCreate, bulkLength)
		for i := 0; i < bulkLength; i++ {
			bulk[i] = client.Project.Create().
				SetID(uint32(num + i + 1)).
				SetUID(toCreates[current+i].UID).
				SetTitle(toCreates[current+i].Title).
				SetGoal(toCreates[current+i].Goal).
				SetPrinciple(toCreates[current+i].Principle).
				SetProcessAndMethod(toCreates[current+i].ProcessAndMethod).
				SetStep(toCreates[current+i].Step).
				SetResultAndConclusion(toCreates[current+i].ResultAndConclusion).
				SetRequirement(toCreates[current+i].Requirement)
		}
		err = client.Project.CreateBulk(bulk...).Exec(ctx)
		if err != nil {
			return fmt.Errorf("create project failed: %w", err)
		}
	}

	return nil
}
