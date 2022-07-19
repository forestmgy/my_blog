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
