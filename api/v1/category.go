package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_blog/model"
	"my_blog/utils/errmsg"
	"net/http"
	"strconv"
)

var scode int

//添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	scode = model.CheckCategoryName(data.Name)
	fmt.Println("状态码", scode)
	if scode == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	fmt.Println("状态码", scode)
	//给前端返回的数据
	c.JSON(http.StatusOK, gin.H{
		"code": scode,
		"data": data,
	})
}

// todo 查询分类下的所有文章

//查询分类列表
//func GetAllCategory(c *gin.Context) {
//	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
//	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
//
//	if pageNum == 0 {
//		pageNum = 1
//	}
//	data, total := model.GetCategory(pageSize, pageNum)
//	code = errmsg.SUCCESS
//	//给前端返回的数据
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"data":    data,
//		"total":   total,
//		"message": errmsg.GetErrMsg(code),
//	})
//}

//编辑分类名
func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	scode = model.CheckCategoryName(data.Name) //这里面并不是在找这个分类是否存在，其实我们是想知道这个分类名是不是被人用过了
	if scode == errmsg.SUCCESS {
		scode = model.EditCategory(id, &data)
	}
	//if code == errmsg.ERROR_USERNAME_DUPLICATED {
	//	c.Abo
	//}
	c.JSON(http.StatusOK, gin.H{
		"status":  scode,
		"message": errmsg.GetErrMsg(code),
	})

}

//删除分类
func DelCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  scode,
		"message": errmsg.GetErrMsg(code),
	})
}
