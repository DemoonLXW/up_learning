package service_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func CreateTestApplicantFacade() (*service.ApplicantFacade, error) {
	os.Chdir("../")
	// os.Setenv("DB_CONFIG", "./database.config.json")
	s, err := injection.ProvideService()
	if err != nil {
		return nil, err
	}

	return s.ApplicantFa, nil
}

func TestCreateProject(t *testing.T) {
	faca, err := CreateTestApplicantFacade()
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
		return faca.CreateProject(ctx, tx.Client(), adds)
	})
	// err = serv.CreateProject(nil, nil, adds)
	assert.Nil(t, err)
}

func TestDeleteProject(t *testing.T) {
	faca, err := CreateTestApplicantFacade()
	assert.Nil(t, err)

	ctx := context.Background()
	client := faca.DB

	ids := []uint32{2, 4}
	deletes, err := client.Project.Query().Where(
		project.IDIn(ids...),
	).All(ctx)
	assert.Nil(t, err)

	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return faca.DeleteProject(ctx, tx.Client(), deletes)
	})

	assert.Nil(t, err)

}

func TestUpdateProject(t *testing.T) {
	faca, err := CreateTestApplicantFacade()
	assert.Nil(t, err)

	id := uint32(4)
	// title := "update title2"
	goal := "update goal4"
	step := "update step4"
	requirement := "update requirement4"
	file1 := entity.ToAddFile{
		UID:  4,
		Name: "testupdate.json",
		Path: "file/uid/testupdate.json",
		Size: 69822,
	}

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
	client := faca.DB

	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return faca.UpdateProject(ctx, tx.Client(), &entity.ToModifyProject{
			ID:            id,
			Goal:          &goal,
			Step:          &step,
			Requirement:   &requirement,
			AddFile:       []*entity.ToAddFile{&file1, &file1},
			DeleteFileIDs: []uint32{9},
		})
	})
	// err = serv.CreateProject(nil, nil, adds)
	assert.Nil(t, err)
}

func TestStartReviewProjectProcess(t *testing.T) {
	faca, err := CreateTestApplicantFacade()
	assert.Nil(t, err)

	userID := uint32(4)
	projectID := uint32(5)

	err = faca.StartReviewProjectProcess(userID, projectID)
	assert.Nil(t, err)

}

func TestRetrieveReviewProjectRecordByProjectID(t *testing.T) {
	faca, err := CreateTestApplicantFacade()
	assert.Nil(t, err)

	pageSize := uint(5)
	current := uint(1)
	sort := "startTime"
	order := false

	projectID := uint32(5)

	m, err := faca.RetrieveReviewProjectRecordByProjectID(&entity.SearchReviewProjectRecord{
		PageSize:  &pageSize,
		Current:   &current,
		Sort:      sort,
		Order:     &order,
		ProjectID: &projectID,
	})
	assert.Nil(t, err)

	for _, v := range m.Data {
		fmt.Print(v.ID, v.StartTime, v.EndTime)
		for _, vv := range v.Variables {
			fmt.Print(vv.Name, vv.Value)
		}
		fmt.Println()
	}
	fmt.Println(m.Total)
}

// func TestFindReviewProjectRecordDetailById(t *testing.T) {
// 	faca, err := CreateTestApplicantFacade()
// 	assert.Nil(t, err)

// 	id := "2eb5fff3-83b6-11ee-8a6b-b66921bdcecd"

// 	pm, tm, err := faca.FindReviewProjectRecordDetailById(id)
// 	assert.Nil(t, err)

// 	for _, v := range pm["data"].([]interface{}) {
// 		mm := v.(map[string]interface{})
// 		fmt.Printf("%v\n", mm["variables"])
// 	}

// 	for _, v := range tm["data"].([]interface{}) {
// 		mm := v.(map[string]interface{})
// 		fmt.Printf("%v\n", mm["name"])
// 	}
// }
