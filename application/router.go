package application

import (
	"github.com/DemoonLXW/up_learning/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine, controller *controller.Controller) *gin.Engine {

	app.POST("/login", controller.User.Login)

	return app
}
