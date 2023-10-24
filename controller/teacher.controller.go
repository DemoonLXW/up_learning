package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	TeacherFa *service.TeacherFacade
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
