package service

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/entity"
)

type TeacherService struct {
	DB *ent.Client
}

func (serv *TeacherService) RetrieveProjectWithFile(ctx context.Context, client *ent.Client, search *entity.SearchProject) ([]*ent.Project, error) {
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
		}).
		WithAttachments(func(fq *ent.FileQuery) {
			fq.Where(func(s *sql.Selector) {
				s.Where(sql.IsNull(file.FieldDeletedTime))
			})
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

func (serv *TeacherService) FindOneProjectWithFileById(ctx context.Context, client *ent.Client, id uint32) (*ent.Project, error) {
	res, err := client.Project.Query().Where(project.And(
		project.IDEQ(id),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(project.FieldDeletedTime))
		},
	)).
		WithAttachments(func(fq *ent.FileQuery) {
			fq.Where(func(s *sql.Selector) {
				s.Where(sql.IsNull(file.FieldDeletedTime))
			})
		}).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("project with file find one by id failed: %w", err)
	}
	return res, nil
}
