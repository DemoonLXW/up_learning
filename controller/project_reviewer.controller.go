package controller

import (
	"context"
	"math"
	"net/http"
	"time"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

type ProjectReviewerController struct {
	ProjectReviewerFa *service.ProjectReviewerFacade
	Applicant         *service.ApplicantService
}

func (cont *ProjectReviewerController) GetTaskListOfPlatformReviewer(c *gin.Context) {
	var search entity.SearchReviewProjectTask
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tpl, err := cont.ProjectReviewerFa.RetrieveReviewProjectTaskOfPlatformReviewer(&search)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	client := cont.Applicant.DB

	tasks := make([]entity.RetrievedReviewProjectTask, 0)
	for _, v := range tpl.Data {
		id := v.ID
		name := v.Name
		createTime := v.CreateTime
		var dueDate *time.Time
		var projectID *uint32
		for _, vv := range v.Variables {
			switch vv.Name {
			case "dueDate":
				t, _ := time.Parse(time.RFC3339, vv.Value.(string))
				dueDate = &t
			case "projectID":
				value := vv.Value.(float64)
				pid := uint32(value)
				projectID = &pid
			}
		}
		if projectID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "project id of task not found"})
			return
		}
		p, err := cont.Applicant.FindOneProjectWithFileAndUserById(ctx, client, *projectID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, entity.RetrievedReviewProjectTask{
			ID:         &id,
			Name:       &name,
			CreateTime: createTime,
			DueDate:    dueDate,
			Project: &entity.RetrievedProject{
				ID:    &p.ID,
				Title: &p.Title,
				Applicant: &entity.RetrievedUser{
					ID:       &p.Edges.User.ID,
					Account:  &p.Edges.User.Account,
					Username: &p.Edges.User.Username,
				},
			},
		})
	}

	var data entity.RetrievedListData
	data.Record = tasks
	total := tpl.Total
	data.Total = total
	if search.Current != nil && search.PageSize != nil {
		data.IsPrevious = *search.Current > 1
		data.IsNext = *search.Current < uint(math.Ceil(float64(total)/float64(*search.PageSize)))
	}

	var res entity.Result
	res.Message = "Get Review Project Task List Of Platform Reviewer Successfully"
	res.Data = data
	c.JSON(http.StatusOK, res)
}
