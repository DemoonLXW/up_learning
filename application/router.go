package application

import (
	"github.com/DemoonLXW/up_learning/controller"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine, controllers *controller.Controllers) *gin.Engine {

	userService := controllers.User.Services.User

	app.POST("/login", controllers.User.Login)
	// app.POST("/autologin", controllers.User.AutoLogin)
	app.POST("/logout", controllers.User.Logout)

	auth := app.Group("/")
	auth.Use(Auth(controllers.User.Services.User))
	{
		auth.POST("/autologin", controllers.User.AutoLogin)

		role := auth.Group("/role")
		{
			role.GET("/getlist", Check(userService, []string{entity.RetrieveRole}), controllers.Management.GetRoleList)
			role.POST("/add", Check(userService, []string{entity.AddRole}), controllers.Management.AddARole)
			role.POST("/modify", Check(userService, []string{entity.ModifyRole}), controllers.Management.ModifyARole)
			role.GET("/get/:id", Check(userService, []string{entity.RetrieveRole}), controllers.Management.GetARoleById)
			role.POST("/remove", Check(userService, []string{entity.RemoveRole}), controllers.Management.RemoveRolesByIds)
			role.GET("/get/permissions/:id", Check(userService, []string{entity.RetrieveRole, entity.RetrievePermission}), controllers.Management.GetPermissionsByRoleId)
		}

		permission := auth.Group("/permission")
		{
			permission.GET("/getlist", Check(userService, []string{entity.RetrievePermission}), controllers.Management.GetPermissionList)
			permission.POST("/add", Check(userService, []string{entity.AddPermission}), controllers.Management.AddAPermission)
			permission.POST("/modify", Check(userService, []string{entity.ModifyPermission}), controllers.Management.ModifyAPermission)
			permission.GET("/get/:id", Check(userService, []string{entity.RetrievePermission}), controllers.Management.GetAPermissionById)
			permission.POST("/remove", Check(userService, []string{entity.RemovePermission}), controllers.Management.RemovePermissionsByIds)
		}

	}

	return app
}
