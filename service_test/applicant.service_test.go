package service_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func CreateTestTeacherService() (*service.ApplicantService, error) {
	os.Chdir("../")
	os.Setenv("DB_CONFIG", "./database.config.json")
	serv := new(service.ApplicantService)
	db, err := injection.ProvideDataBase()
	if err != nil {
		return nil, err
	}
	serv.DB = db

	return serv, nil
}

func TestRetrieveProject(t *testing.T) {
	serv, err := CreateTestTeacherService()
	assert.Nil(t, err)

	ctx := context.Background()
	client := serv.DB

	current := uint(4)
	pageSize := uint(1)
	// var status uint8 = 0
	// order := true
	// disabled := true
	projects, err := serv.RetrieveProjectWithFileAndUserByUserID(ctx, client, &entity.SearchProject{
		// Order: &order,
		// Sort:  sort,
		Current:  &current,
		PageSize: &pageSize,
		// ReviewStatus: &status,
		// Like:         "e2",
		// IsDisabled:   &disabled,
	}, uint32(4))
	assert.Nil(t, err)
	for _, v := range projects {

		fmt.Println(v)
		if v.Edges.Attachments != nil {
			for _, f := range v.Edges.Attachments {
				fmt.Print(f)
			}
		}
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
	total, err := serv.GetTotalRetrievedProjectsByUserID(ctx, client, &entity.SearchProject{
		// Order: &order,
		// Sort:  sort,
		// Current:      &current,
		// PageSize:     &pageSize,
		ReviewStatus: &status,
		Like:         "title",
		// IsDisabled:   &disabled,
	}, uint32(4))
	assert.Nil(t, err)

	assert.Equal(t, 2, total)
}

func TestFindOneProjectWithFileById(t *testing.T) {
	serv, err := CreateTestTeacherService()
	assert.Nil(t, err)

	ctx := context.Background()
	client := serv.DB

	id := uint32(1)
	p, err := serv.FindOneProjectWithFileAndUserById(ctx, client, id)
	assert.Nil(t, err)

	fmt.Println(p)
	for _, v := range p.Edges.Attachments {
		fmt.Println(v.Name)
	}
	if p.Edges.User != nil {
		fmt.Println(p.Edges.User)
	}
}
