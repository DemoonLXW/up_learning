package application

import (
	"github.com/DemoonLXW/up_learning/controller"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine, controllers *controller.Controllers) *gin.Engine {

	authService := controllers.Auth.Services.Auth

	app.POST("/login", controllers.Auth.Login)
	// app.POST("/autologin", controllers.Auth.AutoLogin)
	app.POST("/logout", controllers.Auth.Logout)

	auth := app.Group("/")
	auth.Use(Auth(controllers.Auth.Services.Auth))
	{
		auth.POST("/autologin", controllers.Auth.AutoLogin)

		role := auth.Group("/role")
		{
			role.GET("/getlist", Check(authService, []string{entity.GetRoleList}), controllers.Management.GetRoleList)
			role.POST("/add", Check(authService, []string{entity.AddRole}), controllers.Management.AddRole)
			role.POST("/modify", Check(authService, []string{entity.ModifyARole}), controllers.Management.ModifyARole)
			role.GET("/get/:id", Check(authService, []string{entity.GetARoleById}), controllers.Management.GetARoleById)
			role.POST("/remove", Check(authService, []string{entity.RemoveRolesByIds}), controllers.Management.RemoveRolesByIds)
			role.GET("/:id/get/permissions", Check(authService, []string{entity.GetPermissionsByRoleId}), controllers.Management.GetPermissionsByRoleId)
			role.POST("/modify/permissions", Check(authService, []string{entity.ModifyPermissionsForRoles}), controllers.Management.ModifyPermissionsForRoles)
		}

		permission := auth.Group("/permission")
		{
			permission.GET("/getlist", Check(authService, []string{entity.GetPermissionList}), controllers.Management.GetPermissionList)
			// permission.POST("/add", Check(authService, []string{entity.AddPermission}), controllers.Management.AddPermission)
			permission.POST("/modify", Check(authService, []string{entity.ModifyAPermission}), controllers.Management.ModifyAPermission)
			permission.GET("/get/:id", Check(authService, []string{entity.GetAPermissionById}), controllers.Management.GetAPermissionById)
			// permission.POST("/remove", Check(authService, []string{entity.RemovePermissionsByIds}), controllers.Management.RemovePermissionsByIds)
		}

		user := auth.Group("/user")
		{
			user.GET("/getlist", Check(authService, []string{entity.GetUserList}), controllers.Management.GetUserList)

		}

		// school := auth.Group("/school")
		// {
		// 	school.POST("/import", Check(authService, []string{entity.ImportSchool}), controllers.Management.ImportSchool)
		// 	school.GET("/getlist", Check(authService, []string{entity.GetSchoolList}), controllers.Management.GetSchoolList)
		// 	school.GET("/getsample", Check(authService, []string{entity.GetSampleOfSchoolImport}), controllers.Management.GetSampleOfSchoolImport)
		// }
		teacher := auth.Group("/teacher")
		{
			teacher.GET("/getlist", Check(authService, []string{entity.GetTeacherList}), controllers.Management.GetTeacherList)
			teacher.POST("/import", Check(authService, []string{entity.ImportTeacherByCollegeID}), controllers.Management.ImportTeacherByCollegeID)
		}

		student := auth.Group("/student")
		{
			student.GET("/getlist", Check(authService, []string{entity.GetStudentList}), controllers.Management.GetStudentList)
			student.POST("/import", Check(authService, []string{entity.ImportStudentByClassID}), controllers.Management.ImportStudentByClassID)
		}

		college := auth.Group("/college")
		{
			college.POST("/import", Check(authService, []string{entity.ImportCollege}), controllers.Management.ImportCollege)
			college.GET("/getlist", Check(authService, []string{entity.GetCollegeList}), controllers.Management.GetCollegeList)
			college.GET("/get", Check(authService, []string{entity.GetColleges}), controllers.Management.GetColleges)
		}

		major := auth.Group("/major")
		{
			major.POST("/import", Check(authService, []string{entity.ImportMajorByCollegeID}), controllers.Management.ImportMajorByCollegeID)
			major.GET("/getlist", Check(authService, []string{entity.GetMajorList}), controllers.Management.GetMajorList)
		}

		class := auth.Group("/class")
		{
			class.POST("/import", Check(authService, []string{entity.ImportClassByMajorID}), controllers.Management.ImportClassByMajorID)
			class.GET("/getlist", Check(authService, []string{entity.GetClassList}), controllers.Management.GetClassList)
		}

		file := auth.Group("/file")
		{
			file.GET("/getsample", Check(authService, []string{entity.GetSampleOfImport}), controllers.Management.GetSampleOfImport)
		}

	}

	return app
}
