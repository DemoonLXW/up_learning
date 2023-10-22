package facade_test

import (
	"context"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/facade"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func CreateTestFacade() (*facade.Facades, error) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	faca, err := injection.ProvideFacade()
	if err != nil {
		return nil, err
	}

	return faca, nil
}

func TestCreateProject(t *testing.T) {
	faca, err := CreateTestFacade()
	assert.Nil(t, err)

	project1 := entity.ToAddProject{
		UID:                 5,
		Title:               "title2",
		Goal:                "goal2",
		Principle:           "principle2",
		ProcessAndMethod:    "process and method 2",
		Step:                "step2",
		ResultAndConclusion: "result and conclusion 2",
		Requirement:         "requirement2",
		Attachments: []*entity.ToAddFile{
			{
				UID:  5,
				Name: "test5.file",
				Path: "file/uid/test5.file",
				Size: 55,
			},
			{
				UID:  5,
				Name: "test6.file",
				Path: "file/uid/test6.file",
				Size: 6,
			},
		},
	}

	project2 := entity.ToAddProject{
		UID:                 4,
		Title:               "title5",
		Goal:                "goal5",
		Principle:           "principle5",
		ProcessAndMethod:    "process and method 5",
		Step:                "step5",
		ResultAndConclusion: "result and conclusion 5",
		Requirement:         "requirement5",
		Attachments: []*entity.ToAddFile{
			{
				UID:  4,
				Name: "test3.jpg",
				Path: "file/uid/test3.jpg",
				Size: 336,
			},
		},
	}

	adds := []*entity.ToAddProject{&project2, &project1}

	ctx := context.Background()
	client := faca.DB

	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return faca.Teacher.CreateProject(ctx, tx.Client(), adds)
	})
	// err = serv.CreateProject(nil, nil, adds)
	assert.Nil(t, err)
}

func TestDeleteProject(t *testing.T) {
	faca, err := CreateTestFacade()
	assert.Nil(t, err)

	ctx := context.Background()
	client := faca.DB

	ids := []uint32{2, 4}

	projects, err := client.Project.Query().
		Where(project.IDIn(ids...)).
		WithAttachments().
		All(ctx)
	assert.Nil(t, err)

	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return faca.Teacher.DeleteProject(ctx, tx.Client(), projects)
	})

	assert.Nil(t, err)

}
