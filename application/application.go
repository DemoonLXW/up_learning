package application

import (
	"io"
	"os"

	"github.com/DemoonLXW/up_learning/controller"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ApplicationProvider = wire.NewSet(
	SetupApplication,
)

func SetupApplication(controller *controller.Controller) *gin.Engine {
	filepath, ok := os.LookupEnv("GIN_LOG")
	if !ok {
		filepath = "./gin.log"
	}
	f, _ := os.Create(filepath)
	gin.DefaultWriter = io.MultiWriter(f)
	defer f.Close()

	app := gin.New()

	middlewared_app := SetupMiddleware(app)

	routed_app := SetupRouter(middlewared_app, controller)

	return routed_app
}
