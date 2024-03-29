package service

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/database/ent/student"
	"github.com/DemoonLXW/up_learning/database/ent/teacher"
	"github.com/DemoonLXW/up_learning/database/ent/user"
	"github.com/DemoonLXW/up_learning/entity"
)

type ApplicantService struct {
	DB *ent.Client
}

func (serv *ApplicantService) RetrieveProjectWithFileAndUserByUserID(ctx context.Context, client *ent.Client, search *entity.SearchProject, userID uint32) ([]*ent.Project, error) {
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
				project.UIDEQ(userID),
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
		}).
		WithUser(func(uq *ent.UserQuery) {
			uq.Where(func(s *sql.Selector) {
				s.Where(sql.IsNull(user.FieldDeletedTime))
			})
		})

	if search.PageSize != nil && search.Current != nil {
		offset := int((*search.Current - 1) * (*search.PageSize))
		query.Limit(int(*search.PageSize)).Offset(offset)
	}
	projects, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve project query failed: %w", err)
	}

	return projects, nil
}

func (serv *ApplicantService) GetTotalRetrievedProjectsByUserID(ctx context.Context, client *ent.Client, search *entity.SearchProject, userID uint32) (int, error) {
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
				project.UIDEQ(userID),
			),
		).Count(ctx)

	if err != nil {
		return -1, fmt.Errorf("get retrieved total projects query failed: %w", err)
	}

	return total, nil
}

func (serv *ApplicantService) FindOneProjectWithFileAndUserById(ctx context.Context, client *ent.Client, id uint32) (*ent.Project, error) {
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
		}).
		WithUser(func(uq *ent.UserQuery) {
			uq.Where(func(s *sql.Selector) {
				s.Where(sql.IsNull(user.FieldDeletedTime))
			})
		}).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("project with file find one by id failed: %w", err)
	}
	return res, nil
}

func (serv *ApplicantService) FindUserWithTeacherOrStudentById(ctx context.Context, client *ent.Client, userID uint32) (*ent.User, error) {
	if ctx == nil || client == nil {
		return nil, fmt.Errorf("context or client is nil")
	}

	user, err := client.User.Query().Where(
		user.IDEQ(userID),
		func(s *sql.Selector) {
			s.Where(sql.IsNull(user.FieldDeletedTime))
		},
	).WithStudent(func(sq *ent.StudentQuery) {
		sq.Where(func(s *sql.Selector) {
			s.Where(sql.IsNull(student.FieldDeletedTime))
		})
	}).WithTeacher(func(tq *ent.TeacherQuery) {
		tq.Where(func(s *sql.Selector) {
			s.Where(sql.IsNull(teacher.FieldDeletedTime))
		})
	}).First(ctx)

	if err != nil {
		return nil, fmt.Errorf("find user with teacher or student by id failed: %w", err)
	}

	return user, nil
}
