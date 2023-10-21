package service

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/entity"
)

type TeacherService struct {
	DB *ent.Client
}

func (serv *TeacherService) CreateProject(ctx context.Context, client *ent.Client, toCreates []*entity.ToAddProject) error {
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

func (serv *TeacherService) UpdateProject(ctx context.Context, client *ent.Client, toUpdate *entity.ToModifyProject) error {
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
	if toUpdate.IsDisabled != nil {
		updater.SetIsDisabled(*toUpdate.IsDisabled)
	}
	updater.SetModifiedTime(time.Now())

	num, err := updater.Save(ctx)
	if err != nil {
		return fmt.Errorf("update project failed: %w", err)
	}
	if num == 0 {
		return fmt.Errorf("update project affect 0 row")
	}

	return nil
}

func (serv *TeacherService) RetrieveProject(ctx context.Context, client *ent.Client, search *entity.SearchProject) ([]*ent.Project, error) {
	if ctx == nil || client == nil {
		return nil, fmt.Errorf("context or client is nil")
	}

	if search == nil {
		return nil, fmt.Errorf("retrieve project failed: search is nil")
	}

	query := client.Project.Query().
		Where(
			project.And(
				project.Or(
					project.TitleContains(search.Like),
				),
				func(s *sql.Selector) {
					s.Where(sql.IsNull(project.FieldDeletedTime))
					if search.IsDisabled != nil {
						s.Where(sql.EQ(project.FieldIsDisabled, *search.IsDisabled))
					}
					if search.ReviewStatus != nil {
						s.Where(sql.EQ(project.FieldReviewStatus, *search.ReviewStatus))
					}
				},
			),
		).
		Order(func(s *sql.Selector) {
			isSorted := search.Sort != "" && (search.Sort == project.FieldID || search.Sort == project.FieldUID || search.Sort == project.FieldTitle || search.Sort == project.FieldIsDisabled)
			if isSorted && search.Order != nil {
				if *search.Order {
					s.OrderBy(sql.Desc(search.Sort))
				} else {
					s.OrderBy(sql.Asc(search.Sort))
				}
			}
		})

	if search.PageSize != nil && search.Current != nil {
		offset := (*search.Current - 1) * (*search.PageSize)
		query.Limit(*search.PageSize).Offset(offset)
	}
	projects, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve project query failed: %w", err)
	}

	return projects, nil
}

func (serv *TeacherService) GetTotalRetrievedProjects(ctx context.Context, client *ent.Client, search *entity.SearchProject) (int, error) {
	if ctx == nil || client == nil {
		return -1, fmt.Errorf("context or client is nil")
	}

	total, err := client.Project.Query().
		Where(
			project.And(
				project.Or(
					project.TitleContains(search.Like),
				),
				func(s *sql.Selector) {
					s.Where(sql.IsNull(project.FieldDeletedTime))
					if search.IsDisabled != nil {
						s.Where(sql.EQ(project.FieldIsDisabled, *search.IsDisabled))
					}
					if search.ReviewStatus != nil {
						s.Where(sql.EQ(project.FieldReviewStatus, *search.ReviewStatus))
					}
				},
			),
		).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total projects query failed: %w", err)
	}

	return total, nil
}
