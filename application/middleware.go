package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

func SetupMiddleware(app *gin.Engine) *gin.Engine {
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	// app.Use(cors.New(cors.Config{
	// 	AllowMethods:     []string{"POST", "GET", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return strings.Contains(origin, "localhost") || strings.Contains(origin, "127.0.0.1")
	// 	},
	// }))
	// app.Use(cors.Default())
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

func Auth(serv *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := c.Cookie("uid")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		token, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.ParseUint(uid, 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		new_token, err := serv.CheckCredential(uint32(id), token)
		if err != nil {
			err_string := err.Error()
			switch {
			case strings.Contains(err_string, "token does not exist") || strings.Contains(err_string, "credential is expired"):
				{
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
					return
				}
			default:
				{
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
					return
				}
			}
		}

		domain := c.Value("domain").(string)
		c.SetCookie("uid", uid, 9000, "/", domain, false, true)
		c.SetCookie("token", new_token, 9000, "/", domain, false, true)

		c.Next()

	}
}
