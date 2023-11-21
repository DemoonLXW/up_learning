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
