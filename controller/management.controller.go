package controller

import (
	"math"
	"net/http"
	"os"

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
	if search.Current != nil && search.PageSize != nil {
		data.IsPrevious = *search.Current > 1
		data.IsNext = *search.Current < int(math.Ceil(float64(total)/float64(*search.PageSize)))
	}

	var res entity.Result
	res.Message = "Get List of Roles Successfully"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) AddRole(c *gin.Context) {

	var roles []entity.ToAddRole
	if err := c.ShouldBindJSON(&roles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rs := make([]*entity.ToAddRole, len(roles))
	for i := range roles {
		rs[i] = &roles[i]
	}

	err := cont.Services.Management.CreateRole(rs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Add Role Successfully"
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
	var role entity.RetrievedDetailedRole
	if err := c.ShouldBindUri(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, err := cont.Services.Management.FindOneRoleById(role.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.Name = r.Name
	role.Description = r.Description
	role.IsDisabled = r.IsDisabled
	role.CreatedTime = r.CreatedTime
	role.ModifiedTime = r.ModifiedTime

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
	if search.Current != nil && search.PageSize != nil {
		data.IsPrevious = *search.Current > 1
		data.IsNext = *search.Current < int(math.Ceil(float64(total)/float64(*search.PageSize)))
	}

	var res entity.Result
	res.Message = "Get List of Permissions Successfully"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) AddPermission(c *gin.Context) {

	var permissions []entity.ToAddPermission
	if err := c.ShouldBindJSON(&permissions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ps := make([]*entity.ToAddPermission, len(permissions))
	for i := range permissions {
		ps[i] = &permissions[i]
	}

	err := cont.Services.Management.CreatePermission(ps)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Add Permission Successfully"
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
	var permission entity.RetrievedDetailedPermission
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
	permission.CreatedTime = p.CreatedTime
	permission.ModifiedTime = p.ModifiedTime

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

func (cont *ManagementController) GetPermissionsByRoleId(c *gin.Context) {
	var role entity.RetrievedRole
	if err := c.ShouldBindUri(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ps, err := cont.Services.Management.FindPermissionsByRoleIds([]uint8{role.ID})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permissions := make([]entity.RetrievedPermission, 0, len(ps))
	for _, v := range ps {
		if !v.IsDisabled {
			permissions = append(permissions, entity.RetrievedPermission{
				ID:          v.ID,
				Action:      v.Action,
				Description: v.Description,
				IsDisabled:  v.IsDisabled,
			})
		}
		// permissions[i].ID = v.ID
		// permissions[i].Action = v.Action
		// permissions[i].Description = v.Description
		// permissions[i].IsDisabled = v.IsDisabled
	}

	var res entity.Result
	res.Message = "Get Permissions By Role Id Successfully"
	res.Data = permissions
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) ModifyPermissionsForRoles(c *gin.Context) {
	var params entity.ToModifyPermissionsOfRoles
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cont.Services.Management.UpdatePermissionForRole(params.RIDs, params.PIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Modify Permissions For Roles Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) GetUserList(c *gin.Context) {

	var search entity.Search
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	us, err := cont.Services.Management.RetrieveUser(search.Current, search.PageSize, search.Like, search.Sort, search.Order, search.IsDisabled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total, err := cont.Services.Management.GetTotalRetrievedUsers(search.Like, search.IsDisabled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users := make([]entity.RetrievedUser, len(us))
	for i, v := range us {
		users[i].ID = v.ID
		users[i].Account = v.Account
		users[i].Username = v.Username
		if v.Email != nil {
			users[i].Email = *v.Email
		}
		if v.Phone != nil {
			users[i].Phone = *v.Phone
		}
		users[i].IsDisabled = v.IsDisabled
	}

	var data entity.RetrievedListData
	data.Record = users
	data.Total = total
	if search.Current != nil && search.PageSize != nil {
		data.IsPrevious = *search.Current > 1
		data.IsNext = *search.Current < int(math.Ceil(float64(total)/float64(*search.PageSize)))
	}

	var res entity.Result
	res.Message = "Get List of Users Successfully"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) ImportSchool(c *gin.Context) {
	fh, err := c.FormFile("import")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f, err := cont.Services.Management.SaveImportedFile(fh, wd+"/temp/importSchool", "school")
	defer os.Remove(f.Name())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := cont.Services.Management.ReadSchoolsFromFile(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cont.Services.Management.CreateSchool(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Import Schools Successfully"
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) GetSchoolList(c *gin.Context) {

	var search entity.Search
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ss, err := cont.Services.Management.RetrieveSchool(search.Current, search.PageSize, search.Like, search.Sort, search.Order, search.IsDisabled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total, err := cont.Services.Management.GetTotalRetrievedSchools(search.Like, search.IsDisabled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schools := make([]entity.RetrievedSchool, len(ss))
	for i, v := range ss {
		schools[i].ID = v.ID
		schools[i].Code = v.Code
		schools[i].Name = v.Name
		schools[i].CompetentDepartment = v.CompetentDepartment
		schools[i].Location = v.Location
		schools[i].Remark = v.Remark
		schools[i].IsDisabled = v.IsDisabled
		switch v.EducationLevel {
		case 0:
			schools[i].EducationLevel = "本科"
		case 1:
			schools[i].EducationLevel = "专科"
		}
	}

	var data entity.RetrievedListData
	data.Record = schools
	data.Total = total
	if search.Current != nil && search.PageSize != nil {
		data.IsPrevious = *search.Current > 1
		data.IsNext = *search.Current < int(math.Ceil(float64(total)/float64(*search.PageSize)))
	}

	var res entity.Result
	res.Message = "Get List of Schools Successfully"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) GetSampleOfSchoolImport(c *gin.Context) {

	f, err := cont.Services.Management.FindOneSampleFileByType("import school")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if f.IsDisabled {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is disabled: " + f.Name})
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.FileAttachment(wd+"/"+f.Path, f.Name)
}

func (cont *ManagementController) GetStudentList(c *gin.Context) {

	var search entity.Search
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ss, err := cont.Services.Management.RetrieveStudentWithClassAndUser(search.Current, search.PageSize, search.Like, search.Sort, search.Order, search.IsDisabled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total, err := cont.Services.Management.GetTotalRetrievedStudents(search.Like, search.IsDisabled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	students := make([]entity.RetrievedStudent, len(ss))
	for i, v := range ss {
		students[i].ID = v.ID
		students[i].StudentID = v.StudentID
		students[i].Name = v.Name
		students[i].IsDisabled = v.IsDisabled
		switch v.Gender {
		case 0:
			students[i].Gender = "男"
		case 1:
			students[i].Gender = "女"
		}
		vClass := v.Edges.Class
		if vClass != nil {
			students[i].Class = &entity.RetrievedClass{
				ID:         vClass.ID,
				Grade:      vClass.Grade,
				Name:       vClass.Name,
				IsDisabled: vClass.IsDisabled,
			}
		}
		vUser := v.Edges.User
		if vUser != nil {
			students[i].User = &entity.RetrievedUser{
				ID:         vUser.ID,
				Account:    vUser.Account,
				Username:   vUser.Username,
				IsDisabled: vUser.IsDisabled,
			}
			if vUser.Email != nil {
				students[i].User.Email = *vUser.Email
			}
			if vUser.Phone != nil {
				students[i].User.Phone = *vUser.Phone

			}
		}
	}

	var data entity.RetrievedListData
	data.Record = students
	data.Total = total
	if search.Current != nil && search.PageSize != nil {
		data.IsPrevious = *search.Current > 1
		data.IsNext = *search.Current < int(math.Ceil(float64(total)/float64(*search.PageSize)))
	}

	var res entity.Result
	res.Message = "Get List of Students Successfully"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func (cont *ManagementController) ImportCollege(c *gin.Context) {
	fh, err := c.FormFile("import")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f, err := cont.Services.Management.SaveImportedFile(fh, wd+"/temp/importCollege", "college")
	defer os.Remove(f.Name())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	co, err := cont.Services.Management.ReadCollegesFromFile(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cont.Services.Management.CreateCollege(co)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res entity.Result
	res.Message = "Import Colleges Successfully"
	c.JSON(http.StatusOK, res)
}
