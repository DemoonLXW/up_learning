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
			role.GET("/getlist", Check(authService, []string{entity.RetrieveRole}), controllers.Management.GetRoleList)
			role.POST("/add", Check(authService, []string{entity.AddRole}), controllers.Management.AddRole)
			role.POST("/modify", Check(authService, []string{entity.ModifyRole}), controllers.Management.ModifyARole)
			role.GET("/get/:id", Check(authService, []string{entity.RetrieveRole}), controllers.Management.GetARoleById)
			role.POST("/remove", Check(authService, []string{entity.RemoveRole}), controllers.Management.RemoveRolesByIds)
			role.GET("/:id/get/permissions", Check(authService, []string{entity.RetrieveRole, entity.RetrievePermission}), controllers.Management.GetPermissionsByRoleId)
			role.POST("/modify/permissions", Check(authService, []string{entity.ModifyRole, entity.RetrievePermission}), controllers.Management.ModifyPermissionsForRoles)
		}

		permission := auth.Group("/permission")
		{
			permission.GET("/getlist", Check(authService, []string{entity.RetrievePermission}), controllers.Management.GetPermissionList)
			permission.POST("/add", Check(authService, []string{entity.AddPermission}), controllers.Management.AddAPermission)
			permission.POST("/modify", Check(authService, []string{entity.ModifyPermission}), controllers.Management.ModifyAPermission)
			permission.GET("/get/:id", Check(authService, []string{entity.RetrievePermission}), controllers.Management.GetAPermissionById)
			permission.POST("/remove", Check(authService, []string{entity.RemovePermission}), controllers.Management.RemovePermissionsByIds)
		}

		user := auth.Group("/user")
		{
			user.GET("/getlist", Check(authService, []string{entity.RetrieveUser}), controllers.Management.GetUserList)

		}

	}

	return app
}
