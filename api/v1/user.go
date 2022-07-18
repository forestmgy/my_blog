package v1

import (
	"github.com/gin-gonic/gin"
	"my_blog/model"
	"my_blog/utils/errmsg"
	"my_blog/utils/validator"
	"net/http"
	"strconv"
)

var code int

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.CheckUser(int(data.ID))
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	//给前端返回的数据
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户

//查询用户列表
func GetAllUser(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	//给前端返回的数据
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑用户
func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUserName(data.Username) //这里面并不是在找这个用户是否存在，其实我们是想知道这个用户名是不是被人用过了
	if code == errmsg.SUCCESS {
		code = model.EditUser(id, &data)
	}
	//if code == errmsg.ERROR_USERNAME_DUPLICATED {
	//	c.Abo
	//}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

//删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
