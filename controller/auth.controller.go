package controller

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Services *service.Services
}

func (cont *AuthController) Login(c *gin.Context) {
	var login entity.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// hex_password := []byte(*login.Password)
	// original_password := make([]byte, hex.DecodedLen(len(hex_password)))

	// if _, err := hex.Decode(original_password, hex_password); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// sha1_password := fmt.Sprintf("%x", sha256.Sum256(original_password))

	lower_pwd := strings.ToLower(*login.Password)
	sha256_pwd := fmt.Sprintf("%x", sha256.Sum256([]byte(lower_pwd)))
	user, token, err := cont.Services.Auth.Login(*login.Account, sha256_pwd)
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
		roles_name[i] = v.Name
		roles_ids[i] = v.ID
	}

	ps, err := cont.Services.Management.FindPermissionsByRoleIds(roles_ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	permissions_action := make([]string, 0, len(ps))
	for _, v := range ps {
		if !v.IsDisabled {
			permissions_action = append(permissions_action, v.Action)
		}
	}

	m, err := cont.Services.Auth.FindMenuByRoleIds(roles_ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var menus []*entity.Menu
	for _, menu := range m {
		if !menu.IsDisabled {
			menus = append(menus, menu.JSONMenu...)
		}
	}

	id := fmt.Sprintf("%d", user.ID)
	domain := c.Value("domain").(string)
	c.SetCookie("uid", id, 9000, "/", domain, false, true)
	c.SetCookie("token", token, 9000, "/", domain, false, true)

	var res entity.Result
	res.Message = "Login Successfully"
	res.Data = map[string]interface{}{"roles": roles_name, "menus": menus, "permissions": permissions_action}

	c.JSON(http.StatusOK, res)
}

func (cont *AuthController) Logout(c *gin.Context) {
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
	err = cont.Services.Auth.Logout(uid, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res entity.Result
	res.Message = "Logout Successfully"

	c.JSON(http.StatusOK, res)

}

func (cont *AuthController) AutoLogin(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(uid, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roles, err := cont.Services.Auth.FindRolesWithMenusAndPermissonsByUserId(uint32(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roles_name := make([]string, 0, len(roles))
	var menus []*entity.Menu
	var permissions_action []string
	for _, r := range roles {
		if !r.IsDisabled {
			roles_name = append(roles_name, r.Name)
			for _, menu := range r.Edges.Menu {
				if !menu.IsDisabled {
					menus = append(menus, menu.JSONMenu...)
				}
			}
			for _, permission := range r.Edges.Permissions {
				if !permission.IsDisabled {
					permissions_action = append(permissions_action, permission.Action)
				}
			}
		}
	}

	var res entity.Result
	res.Message = "Auto Login Successfully"
	res.Data = map[string]interface{}{"roles": roles_name, "menus": menus, "permissions": permissions_action}

	c.JSON(http.StatusOK, res)
}
