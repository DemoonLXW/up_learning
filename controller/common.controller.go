package controller

import (
	"context"
	"net/http"

	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

type CommonController struct {
	Common *service.CommonService
}

func (cont *CommonController) DownloadFile(c *gin.Context) {

	var download struct {
		IDs []uint32 `form:"ids" binding:"gte=1"`
	}

	if err := c.ShouldBindQuery(&download); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	client := cont.Common.DB

	f, err := cont.Common.DownloadFilesByIDs(ctx, client, download.IDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.FileAttachment(*f.Path, *f.Name)

}
