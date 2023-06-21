package controller

import (
	"fmt"
	"net/http"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.Service
}

func (cont *UserController) Login(c *gin.Context) {
	var login entity.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := cont.service.User.Login(*login.Account, *login.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := fmt.Sprintf("%d", user.ID)
	domain := c.Value("domain").(string)
	c.SetCookie("uid", id, 9000, "/", domain, false, true)
	c.SetCookie("token", token, 9000, "/", domain, false, true)

	var res entity.Result
	roles := make([]string, len(user.Edges.Roles))
	for i, v := range user.Edges.Roles {
		roles[i] = *v.Name
	}
	res.Message = "Login Successfully"
	res.Data = map[string]interface{}{"roles": roles}

	c.JSON(http.StatusOK, res)
}
