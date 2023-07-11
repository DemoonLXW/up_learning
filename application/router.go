package application

import (
	"github.com/DemoonLXW/up_learning/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine, controllers *controller.Controllers) *gin.Engine {

	app.POST("/login", controllers.User.Login)
	// app.POST("/autologin", controllers.User.AutoLogin)
	app.POST("/logout", controllers.User.Logout)

	auth := app.Group("/")
	auth.Use(Auth(controllers.User.Services.User))

	auth.POST("/autologin", controllers.User.AutoLogin)

	return app
}
