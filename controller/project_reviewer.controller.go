package controller

import (
	"context"
	"math"
	"net/http"
	"strconv"
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

func (cont *ProjectReviewerController) GetATaskDetailByID(c *gin.Context) {
	type Task struct {
		ID *string `uri:"id" binding:"required"`
	}
	var t Task
	if err := c.ShouldBindUri(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hti, err := cont.ProjectReviewerFa.FindReviewProjectTaskDetailByID(*t.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := entity.RetrievedReviewProjectTask{}
	task.ID = &hti.ID
	task.CreateTime = hti.StartTime
	task.Name = &hti.Name
	var projectID *uint32
	for _, v := range hti.Variables {
		switch v.Name {
		case "dueDate":
			t, _ := time.Parse(time.RFC3339, v.Value.(string))
			task.DueDate = &t
		case "projectID":
			value := v.Value.(float64)
			pid := uint32(value)
			projectID = &pid
		}
	}
	if projectID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "project id of task not found"})
		return
	}
	ctx := context.Background()
	client := cont.Applicant.DB
	p, err := cont.Applicant.FindOneProjectWithFileAndUserById(ctx, client, *projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project := &entity.RetrievedProject{
		ID:                  &p.ID,
		Title:               &p.Title,
		Goal:                &p.Goal,
		Principle:           &p.Principle,
		ProcessAndMethod:    &p.ProcessAndMethod,
		Step:                &p.Step,
		ResultAndConclusion: &p.ResultAndConclusion,
		Requirement:         &p.Requirement,
		CreatedTime:         p.CreatedTime,
	}
	applicant := p.Edges.User
	if applicant != nil {
		project.Applicant = &entity.RetrievedUser{
			ID:       &applicant.ID,
			Account:  &applicant.Account,
			Username: &applicant.Username,
		}
	}
	attachments := p.Edges.Attachments
	for _, v := range attachments {
		project.Attachments = append(project.Attachments, &entity.RetrievedFile{
			ID:   &v.ID,
			Name: &v.Name,
			Size: &v.Size,
		})
	}

	task.Project = project

	var res entity.Result
	res.Message = "Get a Task Detail By ID Successfully"
	res.Data = task
	c.JSON(http.StatusOK, res)
}

func (cont *ProjectReviewerController) ReviewProjectByTaskID(c *gin.Context) {

	uid, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := strconv.ParseUint(uid, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var r entity.Review
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cont.ProjectReviewerFa.CompleteReviewProjectBytaskID(*r.ID, uint32(userID), *r.Action)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Review project by taskID successfully"
	c.JSON(http.StatusOK, res)
}
