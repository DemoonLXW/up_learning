package service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func CreateTestProjectReviewerFacade() (*service.ProjectReviewerFacade, error) {
	os.Chdir("../")
	// os.Setenv("DB_CONFIG", "./database.config.json")
	s, err := injection.ProvideService()
	if err != nil {
		return nil, err
	}

	return s.ProjectReviewerFa, nil
}

func TestRetrieveReviewProjectTaskOfTeacher(t *testing.T) {
	faca, err := CreateTestProjectReviewerFacade()
	assert.Nil(t, err)

	uid := uint32(4)

	pageSize := uint(5)
	current := uint(1)
	sort := "id"
	order := false

	m, err := faca.RetrieveReviewProjectTaskOfTeacher(&entity.SearchReviewProjectTask{
		PageSize: &pageSize,
		Current:  &current,
		Sort:     sort,
		Order:    &order,
	}, uid)
	assert.Nil(t, err)

	for _, v := range m.Data {
		fmt.Print(v.ID, v.CreateTime)
		for _, vv := range v.Variables {
			fmt.Print(vv.Name, vv.Value)
		}
		fmt.Println()
	}
	fmt.Println(m.Total)
}

func TestRetrieveReviewProjectTaskOfTeacherReviewer(t *testing.T) {
	faca, err := CreateTestProjectReviewerFacade()
	assert.Nil(t, err)

	cid := uint8(10)

	pageSize := uint(5)
	current := uint(1)
	sort := "id"
	order := false

	m, err := faca.RetrieveReviewProjectTaskOfTeacherReviewer(&entity.SearchReviewProjectTask{
		PageSize: &pageSize,
		Current:  &current,
		Sort:     sort,
		Order:    &order,
	}, cid)
	assert.Nil(t, err)

	for _, v := range m.Data {
		fmt.Print(v.ID, v.CreateTime)
		for _, vv := range v.Variables {
			fmt.Print(vv.Name, vv.Value)
		}
		fmt.Println()
	}
	fmt.Println(m.Total)
}

func TestRetrieveReviewProjectTaskOfPlatformReviewer(t *testing.T) {
	faca, err := CreateTestProjectReviewerFacade()
	assert.Nil(t, err)

	pageSize := uint(5)
	current := uint(1)
	sort := "id"
	order := false

	m, err := faca.RetrieveReviewProjectTaskOfPlatformReviewer(&entity.SearchReviewProjectTask{
		PageSize: &pageSize,
		Current:  &current,
		Sort:     sort,
		Order:    &order,
	})
	assert.Nil(t, err)

	for _, v := range m.Data {
		fmt.Print(v.ID, v.CreateTime)
		for _, vv := range v.Variables {
			fmt.Print(vv.Name, vv.Value)
		}
		fmt.Println()
	}
	fmt.Println(m.Total)
}

func TestFindReviewProjectTaskDetailByID(t *testing.T) {
	faca, err := CreateTestProjectReviewerFacade()
	assert.Nil(t, err)

	taskID := "a47b42e4-8840-11ee-97b1-0242aec148e6"

	hti, err := faca.FindReviewProjectTaskDetailByID(taskID)
	assert.Nil(t, err)

	fmt.Println(hti.ID, hti.Name)
	for _, v := range hti.Variables {
		fmt.Println(v.Name, v.Value)
	}

}

func TestReviewProjectBytaskID(t *testing.T) {
	faca, err := CreateTestProjectReviewerFacade()
	assert.Nil(t, err)

	taskID := "de5ad097-8887-11ee-a7eb-0242aec148e6"
	reviewerID := uint32(5)
	action := uint8(1)

	err = faca.ReviewProjectBytaskID(taskID, reviewerID, action)
	assert.Nil(t, err)

}
