package application

import (
	"github.com/DemoonLXW/up_learning/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine, controllers *controller.Controllers) *gin.Engine {

	app.POST("/login", controllers.User.Login)

	return app
}
