package controller

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

type CommonController struct {
	Services *service.Services
}

func (cont *CommonController) UploadDocumentImage(c *gin.Context) {
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

	f, err := cont.Services.Common.SaveUploadFile(fh, wd+"/files/images/"+uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domain := c.Value("domain").(string)
	fname := filepath.Base(f.Name())
	url := "http://" + domain + "/images/" + uid + "/" + fname

	var res entity.Result
	res.Data = gin.H{"url": url, "alt": fname, "href": url}
	res.Message = "Save Upload Image Successfully"

	c.JSON(http.StatusOK, res)
}
