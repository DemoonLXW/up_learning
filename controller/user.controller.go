package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Services *service.Services
}

func (cont *UserController) Login(c *gin.Context) {
	var login entity.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hex_password := []byte(*login.Password)
	original_password := make([]byte, hex.DecodedLen(len(hex_password)))

	if _, err := hex.Decode(original_password, hex_password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sha1_password := fmt.Sprintf("%x", sha256.Sum256(original_password))
	user, token, err := cont.Services.User.Login(*login.Account, sha1_password)
	if err != nil {
		err_string := err.Error()
		switch {
		case strings.Contains(err_string, "user not found"):
			{
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
		default:
			{
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

	}

	roles_name := make([]string, len(user.Edges.Roles))
	roles_ids := make([]uint8, len(user.Edges.Roles))
	for i, v := range user.Edges.Roles {
		roles_name[i] = *v.Name
		roles_ids[i] = v.ID
	}

	m, err := cont.Services.User.FindMenuByRoleIds(roles_ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	menus := make([]*entity.Menu, 0)
	for _, menu := range m {
		menus = append(menus, menu.JSONMenu...)
	}

	id := fmt.Sprintf("%d", user.ID)
	domain := c.Value("domain").(string)
	c.SetCookie("uid", id, 9000, "/", domain, false, true)
	c.SetCookie("token", token, 9000, "/", domain, false, true)

	var res entity.Result
	res.Message = "Login Successfully"
	res.Data = map[string]interface{}{"roles": roles_name, "menus": menus}

	c.JSON(http.StatusOK, res)
}

func (cont *UserController) Logout(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = cont.Services.User.Logout(uid, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res entity.Result
	res.Message = "Logout Successfully"

	c.JSON(http.StatusOK, res)

}

func (cont *UserController) AutoLogin(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(uid, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	new_token, err := cont.Services.User.CheckCredential(uint32(id), token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m, err := cont.Services.User.FindMenuByUserId(uint32(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	menus := make([]*entity.Menu, 0)
	for _, menu := range m {
		menus = append(menus, menu.JSONMenu...)
	}

	domain := c.Value("domain").(string)
	c.SetCookie("uid", uid, 9000, "/", domain, false, true)
	c.SetCookie("token", new_token, 9000, "/", domain, false, true)

	var res entity.Result
	res.Message = "Auto Login Successfully"
	res.Data = map[string][]*entity.Menu{"menus": menus}

	c.JSON(http.StatusOK, res)
}
