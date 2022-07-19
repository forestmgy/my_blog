package v1

import (
	"github.com/gin-gonic/gin"
	"my_blog/middleware"
	"my_blog/model"
	"my_blog/utils"
	"my_blog/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var code int
	var token string
	code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

func LoginView(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title":       utils.Title,
		"Description": utils.Description,
		"Logo":        utils.Logo,
		"Navigation":  utils.Navigation,
	})
}
