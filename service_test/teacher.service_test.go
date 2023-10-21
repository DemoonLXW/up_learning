package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func CreateTestTeacherService() (*service.TeacherService, error) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.TeacherService)
	db, err := injection.ProvideDataBase()
	if err != nil {
		return nil, err
	}
	serv.DB = db

	return serv, nil
}

func TestCreateProject(t *testing.T) {
	serv, err := CreateTestTeacherService()
	assert.Nil(t, err)

	project1 := entity.ToAddProject{
		UID:                 2,
		Title:               "title1",
		Goal:                "goal1",
		Principle:           "principle1",
		ProcessAndMethod:    "process and method 1",
		Step:                "step1",
		ResultAndConclusion: "result and conclusion 1",
		Requirement:         "requirement1",
	}

	project2 := entity.ToAddProject{
		UID:                 2,
		Title:               "title2",
		Goal:                "goal2",
		Principle:           "principle2",
		ProcessAndMethod:    "process and method 2",
		Step:                "step2",
		ResultAndConclusion: "result and conclusion 2",
		Requirement:         "requirement2",
	}

	adds := []*entity.ToAddProject{&project2, &project1}

	ctx := context.Background()
	client := serv.DB

	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return serv.CreateProject(ctx, tx.Client(), adds)
	})
	// err = serv.CreateProject(nil, nil, adds)
	assert.Nil(t, err)
}

func TestUpdateProject(t *testing.T) {
	serv, err := CreateTestTeacherService()
	assert.Nil(t, err)

	id := uint32(2)
	// title := "update title2"
	goal := "update goal1"
	step := "update step1"
	requirement := "update requirement1"

	// project := entity.ToModifyProject{
	// 	Title:               "title1",
	// 	Goal:                "goal1",
	// 	Principle:           "principle1",
	// 	ProcessAndMethod:    "process and method 1",
	// 	Step:                "step1",
	// 	ResultAndConclusion: "result and conclusion 1",
	// 	Requirement:         "requirement1",
	// }

	ctx := context.Background()
	client := serv.DB

	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return serv.UpdateProject(ctx, tx.Client(), &entity.ToModifyProject{
			ID:          id,
			Goal:        &goal,
			Step:        &step,
			Requirement: &requirement,
		})
	})
	// err = serv.CreateProject(nil, nil, adds)
	assert.Nil(t, err)
}
