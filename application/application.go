package application

import (
	"io"
	"os"
	"path/filepath"

	"github.com/DemoonLXW/up_learning/controller"
	"github.com/DemoonLXW/up_learning/workflow"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ApplicationProvider = wire.NewSet(
	SetupApplication,
)

func SetupApplication(controllers *controller.Controllers, workflowHelper *workflow.WorkflowHelper) (*gin.Engine, error) {
	fp, ok := os.LookupEnv("GIN_LOG")
	if !ok {
		fp = "./gin.log"
	}

	dir := filepath.Dir(fp)
	if err := os.MkdirAll(dir, 0750); err != nil {
		return nil, err
	}

	f, _ := os.Create(fp)
	gin.DefaultWriter = io.MultiWriter(f)

	app := gin.New()

	app.Static("/images", "./files/images")

	middlewared_app := SetupMiddleware(app)

	routed_app := SetupRouter(middlewared_app, controllers)

	err := workflow.SetupWorkflow(workflowHelper)
	if err != nil {
		return nil, err
	}

	return routed_app, nil
}
