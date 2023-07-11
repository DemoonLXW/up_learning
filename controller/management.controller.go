package controller

import (
	"net/http"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

type ManagementController struct {
	Services *service.Services
}

func (cont *ManagementController) GetAllRoles(c *gin.Context) {

	var search entity.Search
	if err := c.ShouldBindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "test get all roles"
	res.Data = map[string]interface{}{"body": search}
	c.JSON(http.StatusOK, res)
}
