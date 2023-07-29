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
		}

	}

	return app
}
