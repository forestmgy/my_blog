package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_blog/middleware"
	"my_blog/model"
	"my_blog/utils"
	"my_blog/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println(err)
	}
	var token string
	var userinfo model.UserInfo
	//t := c.Request.Header
	//body := c.Request
	code, user := model.CheckLogin(data.Username, data.Password)
	fmt.Println(data)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	userinfo.ID = int(user.ID)
	userinfo.UserName = user.Username
	fmt.Println(userinfo)
	println(token)
	c.JSON(http.StatusOK, gin.H{
		"code":     code,
		"userInfo": userinfo,
		"token":    token,
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
