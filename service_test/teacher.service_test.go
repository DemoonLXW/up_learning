package service_test

import (
	"context"
	"fmt"
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

// func TestCreateProject(t *testing.T) {
// 	serv, err := CreateTestTeacherService()
// 	assert.Nil(t, err)

// 	project4 := entity.ToAddProject{
// 		UID:                 3,
// 		Title:               "title4",
// 		Goal:                "goal4",
// 		Principle:           "principle4",
// 		ProcessAndMethod:    "process and method 4",
// 		Step:                "step4",
// 		ResultAndConclusion: "result and conclusion 4",
// 		Requirement:         "requirement4",
// 	}

// 	project3 := entity.ToAddProject{
// 		UID:                 4,
// 		Title:               "title3",
// 		Goal:                "goal3",
// 		Principle:           "principle3",
// 		ProcessAndMethod:    "process and method 3",
// 		Step:                "step3",
// 		ResultAndConclusion: "result and conclusion 3",
// 		Requirement:         "requirement3",
// 	}

// 	adds := []*entity.ToAddProject{&project4, &project3}

// 	ctx := context.Background()
// 	client := serv.DB

// 	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
// 		return serv.CreateProject(ctx, tx.Client(), adds)
// 	})
// err = serv.CreateProject(nil, nil, adds)
// 	assert.Nil(t, err)
// }

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

func TestRetrieveProject(t *testing.T) {
	serv, err := CreateTestTeacherService()
	assert.Nil(t, err)

	ctx := context.Background()
	client := serv.DB

	current := 2
	pageSize := 1
	// var status uint8 = 0
	// order := true
	// disabled := true
	projects, err := serv.RetrieveProject(ctx, client, &entity.SearchProject{
		// Order: &order,
		// Sort:  sort,
		Current:  &current,
		PageSize: &pageSize,
		// ReviewStatus: &status,
		// Like:         "e2",
		// IsDisabled:   &disabled,
	})
	assert.Nil(t, err)
	for _, v := range projects {

		fmt.Println(v)
	}
	assert.Len(t, projects, 1)
}

func TestGetTotalRetrievedProjects(t *testing.T) {
	serv, err := CreateTestTeacherService()
	assert.Nil(t, err)

	ctx := context.Background()
	client := serv.DB

	// current := 1
	// pageSize := 3
	var status uint8 = 0
	// order := true
	// disabled := false
	total, err := serv.GetTotalRetrievedProjects(ctx, client, &entity.SearchProject{
		// Order: &order,
		// Sort:  sort,
		// Current:      &current,
		// PageSize:     &pageSize,
		ReviewStatus: &status,
		Like:         "title",
		// IsDisabled:   &disabled,
	})
	assert.Nil(t, err)

	assert.Equal(t, 2, total)
}

// func TestDeleteProject(t *testing.T) {
// 	serv, err := CreateTestTeacherService()
// 	assert.Nil(t, err)

// 	ctx := context.Background()
// 	client := serv.DB

// 	projects := []*ent.Project{
// 		{
// 			ID:    1,
// 			Title: "title2",
// 		},
// {
// 	ID:    12,
// 	Title: "title2",
// },
// 	}

// 	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
// 		return serv.DeleteProject(ctx, tx.Client(), projects)
// 	})

// 	assert.Nil(t, err)

// }

func TestFindProjectByIDs(t *testing.T) {
	serv, err := CreateTestTeacherService()
	assert.Nil(t, err)

	ctx := context.Background()
	client := serv.DB

	ids := []uint32{2, 3}

	projects, err := serv.FindProjectByIDs(ctx, client, ids)
	assert.Nil(t, err)
	for i := range projects {
		fmt.Println(projects[i])
	}

}
