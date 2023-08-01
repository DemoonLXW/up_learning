package controller

import (
	"math"
	"net/http"

	"github.com/DemoonLXW/up_learning/database/ent"
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

	roles := make([]entity.RetrievedRole, len(rs))
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

	id, err := cont.Services.Management.FindADeletedRoleID()
	if err != nil {
		if !ent.IsNotFound(err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := cont.Services.Management.CreateRole([]*entity.ToAddRole{&role})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		isDisabled := false
		toCreateRole := &entity.ToModifyRole{ID: id, Name: role.Name, Description: role.Description, IsDisabled: &isDisabled}
		err := cont.Services.Management.UpdateDeletedRole(toCreateRole)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	var res entity.Result
	res.Message = "Add a Role Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) ModifyARole(c *gin.Context) {
	var role entity.ToModifyRole
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cont.Services.Management.UpdateRole(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Modify a Role Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) GetARoleById(c *gin.Context) {
	var roleParam entity.RetrievedRole
	if err := c.ShouldBindUri(&roleParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, err := cont.Services.Management.FindOneRoleById(roleParam.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := entity.RetrievedRole{ID: r.ID, Name: r.Name, Description: r.Description, IsDisabled: r.IsDisabled}
	var res entity.Result
	res.Message = "Get a Role By Id Successfully"
	res.Data = role
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) RemoveRolesByIds(c *gin.Context) {
	var ids entity.ToRemoveRoleIDs
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cont.Services.Management.DeleteRole(ids.IDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Remove Roles By Ids Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) GetPermissionList(c *gin.Context) {
	var search entity.Search
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ps, err := cont.Services.Management.RetrievePermission(search.Current, search.PageSize, search.Like, search.Sort, search.Order, search.IsDisabled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total, err := cont.Services.Management.GetTotalRetrievedPermissions(search.Like, search.IsDisabled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permissions := make([]entity.RetrievedPermission, len(ps))
	for i, v := range ps {
		permissions[i].ID = v.ID
		permissions[i].Action = v.Action
		permissions[i].Description = v.Description
		permissions[i].IsDisabled = v.IsDisabled
	}

	var data entity.RetrievedListData
	data.Record = permissions
	data.Total = total
	data.IsPrevious = search.Current > 1
	data.IsNext = search.Current < int(math.Ceil(float64(total)/float64(search.PageSize)))

	var res entity.Result
	res.Message = "Get List of Permissions Successfully"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) AddAPermission(c *gin.Context) {

	var permission entity.ToAddPermission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := cont.Services.Management.FindADeletedPermissionID()
	if err != nil {
		if !ent.IsNotFound(err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := cont.Services.Management.CreatePermission([]*entity.ToAddPermission{&permission})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		isDisabled := false
		toCreatePermission := &entity.ToModifyPermission{ID: id, Action: permission.Action, Description: permission.Description, IsDisabled: &isDisabled}
		err := cont.Services.Management.UpdateDeletedPermission(toCreatePermission)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	var res entity.Result
	res.Message = "Add a Permission Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) ModifyAPermission(c *gin.Context) {
	var permission entity.ToModifyPermission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cont.Services.Management.UpdatePermission(&permission)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Modify a Permission Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) GetAPermissionById(c *gin.Context) {
	var permission entity.RetrievedPermission
	if err := c.ShouldBindUri(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := cont.Services.Management.FindOnePermissionById(permission.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission.Action = p.Action
	permission.Description = p.Description
	permission.IsDisabled = p.IsDisabled
	var res entity.Result
	res.Message = "Get a Permission By Id Successfully"
	res.Data = permission
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) RemovePermissionsByIds(c *gin.Context) {
	var ids entity.ToRemovePermissionIDs
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cont.Services.Management.DeletePermission(ids.IDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Remove Permissions By Ids Successfully"
	c.JSON(http.StatusOK, res)
}
