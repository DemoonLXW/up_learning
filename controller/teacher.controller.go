package controller

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	TeacherFa *service.TeacherFacade
	Teacher   *service.TeacherService
	Common    *service.CommonService
}

func (cont *TeacherController) AddProject(c *gin.Context) {
	var project entity.ToAddProject
	if err := c.ShouldBind(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	project.UID = uint32(userID)

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dir := "files/project/" + uid

	fhs, ok := form.File["attachments"]
	if ok {
		files := make([]*entity.ToAddFile, len(fhs))
		for i := range fhs {
			addFile, err := cont.Common.SaveUploadFile(fhs[i], dir)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			addFile.UID = uint32(userID)

			files[i] = addFile
		}
		project.Attachments = files
	}

	ctx := context.Background()
	client := cont.TeacherFa.DB
	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return cont.TeacherFa.CreateProject(ctx, tx.Client(), []*entity.ToAddProject{&project})
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Add Project Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *TeacherController) RemoveProjectsByIds(c *gin.Context) {
	var remove entity.ToRemoveProjectIDs
	if err := c.ShouldBindJSON(&remove); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	client := cont.TeacherFa.DB
	err := service.WithTx(ctx, client, func(tx *ent.Tx) error {
		toDeletes, err := tx.Client().Project.Query().Where(project.And(
			project.IDIn(remove.IDs...),
			func(s *sql.Selector) {
				s.Where(sql.IsNull(project.FieldDeletedTime))
			},
			project.ReviewStatusNotIn(uint8(1), uint8(2)),
		)).
			WithAttachments().
			All(ctx)
		if err != nil || len(remove.IDs) != len(toDeletes) {
			return fmt.Errorf("delete project not found projects or project is ineligible failed: %w", err)
		}

		return cont.TeacherFa.DeleteProject(ctx, tx.Client(), toDeletes)
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Remove Projects By Ids Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *TeacherController) ModifyAProject(c *gin.Context) {
	var project entity.ToModifyProject
	if err := c.ShouldBind(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dir := "files/project/" + uid

	fhs, ok := form.File["add_file"]
	if ok {
		files := make([]*entity.ToAddFile, len(fhs))
		for i := range fhs {
			addFile, err := cont.Common.SaveUploadFile(fhs[i], dir)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			addFile.UID = uint32(userID)

			files[i] = addFile
		}
		project.AddFile = files
	}

	ctx := context.Background()
	client := cont.TeacherFa.DB
	// can not change review status here
	project.ReviewStatus = nil
	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return cont.TeacherFa.UpdateProject(ctx, tx.Client(), &project)
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Modify a Project Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *TeacherController) GetAProjectById(c *gin.Context) {
	var project entity.RetrievedProject
	if err := c.ShouldBindUri(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	client := cont.Teacher.DB
	object, err := cont.Teacher.FindOneProjectWithFileAndUserById(ctx, client, *project.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if object.IsDisabled {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the project is disabled"})
		return
	}

	project.Title = &object.Title
	project.Goal = &object.Goal
	project.Principle = &object.Principle
	project.ProcessAndMethod = &object.ProcessAndMethod
	project.Step = &object.Step
	project.ResultAndConclusion = &object.ResultAndConclusion
	project.Requirement = &object.Requirement
	project.ReviewStatus = &object.ReviewStatus
	project.CreatedTime = object.CreatedTime
	project.ModifiedTime = object.ModifiedTime
	user := object.Edges.User
	if user != nil {
		project.User = &entity.RetrievedUser{
			ID:       &user.ID,
			Account:  &user.Account,
			Username: &user.Username,
		}
	}
	attachments := object.Edges.Attachments
	for _, v := range attachments {
		project.Attachments = append(project.Attachments, &entity.RetrievedFile{
			ID:   &v.ID,
			Name: &v.Name,
			Size: &v.Size,
		})
	}

	var res entity.Result
	res.Message = "Get a Project By Id Successfully"
	res.Data = project
	c.JSON(http.StatusOK, res)
}

func (cont *TeacherController) GetProjectListByUserID(c *gin.Context) {
	var search entity.SearchProject
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	ctx := context.Background()
	client := cont.Teacher.DB
	objects, err := cont.Teacher.RetrieveProjectWithFileAndUserByUserID(ctx, client, &search, uint32(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total, err := cont.Teacher.GetTotalRetrievedProjectsByUserID(ctx, client, &search, uint32(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var projects []entity.RetrievedProject
	for _, v := range objects {
		if v.IsDisabled {
			total--
			continue
		}
		var p entity.RetrievedProject
		p.ID = &v.ID
		p.Title = &v.Title
		p.ReviewStatus = &v.ReviewStatus
		projects = append(projects, p)
	}

	var data entity.RetrievedListData
	data.Record = projects
	data.Total = total
	if search.Current != nil && search.PageSize != nil {
		data.IsPrevious = *search.Current > 1
		data.IsNext = *search.Current < uint(math.Ceil(float64(total)/float64(*search.PageSize)))
	}

	var res entity.Result
	res.Message = "Get List of My Projects Successfully"
	res.Data = data
	c.JSON(http.StatusOK, res)
}
