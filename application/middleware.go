package application

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupMiddleware(app *gin.Engine) *gin.Engine {
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(CookieDomain())

	return app
}

func GetDomainConfig() (map[string]interface{}, error) {
	filepath, ok := os.LookupEnv("DOMAIN_CONFIG")
	if !ok {
		filepath = "./domain.config.json"
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("read domain config of database failed: %w", err)
	}

	var config map[string]interface{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config of domain failed: %w", err)
	}
	return config, nil
}

func CookieDomain() gin.HandlerFunc {
	return func(c *gin.Context) {
		config, err := GetDomainConfig()
		if err != nil {
			panic(err.Error())
		}
		domain, ok := config["domain"]
		if !ok {
			panic("read domain in config failed")
		}
		c.Set("domain", domain.(string))

		c.Next()

	}
}
