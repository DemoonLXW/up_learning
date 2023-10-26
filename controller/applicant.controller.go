package controller

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ApplicantController struct {
	ApplicantFa *service.ApplicantFacade
	Applicant   *service.ApplicantService
	Common      *service.CommonService
}

func (cont *ApplicantController) AddProject(c *gin.Context) {
	var project entity.ToAddProject
	if err := c.ShouldBindWith(&project, binding.Form); err != nil {
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
	client := cont.ApplicantFa.DB
	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return cont.ApplicantFa.CreateProject(ctx, tx.Client(), []*entity.ToAddProject{&project})
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Add Project Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *ApplicantController) RemoveProjectsByIds(c *gin.Context) {
	var remove entity.ToRemoveProjectIDs
	if err := c.ShouldBindJSON(&remove); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	client := cont.ApplicantFa.DB
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

		return cont.ApplicantFa.DeleteProject(ctx, tx.Client(), toDeletes)
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Remove Projects By Ids Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *ApplicantController) ModifyAProject(c *gin.Context) {
	var project entity.ToModifyProject
	if err := c.ShouldBindWith(&project, binding.Form); err != nil {
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
	client := cont.ApplicantFa.DB
	// can not change review status here
	project.ReviewStatus = nil
	err = service.WithTx(ctx, client, func(tx *ent.Tx) error {
		return cont.ApplicantFa.UpdateProject(ctx, tx.Client(), &project)
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Modify a Project Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *ApplicantController) GetAProjectById(c *gin.Context) {
	var project entity.RetrievedProject
	if err := c.ShouldBindUri(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	client := cont.Applicant.DB
	object, err := cont.Applicant.FindOneProjectWithFileAndUserById(ctx, client, *project.ID)
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

func (cont *ApplicantController) GetProjectListByUserID(c *gin.Context) {
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
	client := cont.Applicant.DB
	objects, err := cont.Applicant.RetrieveProjectWithFileAndUserByUserID(ctx, client, &search, uint32(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total, err := cont.Applicant.GetTotalRetrievedProjectsByUserID(ctx, client, &search, uint32(userID))
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
	res.Message = "Get List of Projects By UserID Successfully"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (cont *ApplicantController) UploadDocumentImage(c *gin.Context) {
	fh, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f, err := cont.Common.SaveUploadFile(fh, wd+"/files/images/"+uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domain := c.Value("domain").(string)
	fname := filepath.Base(f.Path)
	url := "http://" + domain + "/images/" + uid + "/" + fname

	var res entity.Result
	res.Data = gin.H{"url": url, "alt": fname, "href": url}
	res.Message = "Save Upload Image Successfully"

	c.JSON(http.StatusOK, res)
}
