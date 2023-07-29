package controller

import (
	"math"
	"net/http"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

type ManagementController struct {
	Services *service.Services
}

func (cont *ManagementController) GetRoleList(c *gin.Context) {

	var search entity.Search
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rs, err := cont.Services.Management.RetrieveRole(search.Current, search.PageSize, search.Like, search.Sort, search.Order, search.IsDisabled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total, err := cont.Services.Management.GetTotalRetrievedRoles(search.Like, search.IsDisabled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roles := make([]entity.RetrievedRoles, len(rs))
	for i, v := range rs {
		roles[i].ID = v.ID
		roles[i].Name = v.Name
		roles[i].Description = v.Description
		roles[i].IsDisabled = v.IsDisabled
	}

	var data entity.RetrievedListData
	data.Record = roles
	data.Total = total
	data.IsPrevious = search.Current > 1
	data.IsNext = search.Current < int(math.Ceil(float64(total)/float64(search.PageSize)))

	var res entity.Result
	res.Message = "Get List of Roles Successfully"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) AddARole(c *gin.Context) {

	var role entity.ToAddRole
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cont.Services.Management.CreateRole([]*entity.ToAddRole{&role})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Add a Role Successfully"
	c.JSON(http.StatusOK, res)
}
