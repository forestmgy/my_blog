package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"my_blog/model"
	"my_blog/utils"
	"my_blog/utils/errmsg"
	"net/http"
	"strconv"
	"strings"
)

//添加文章
func AddArt(c *gin.Context) {
	var data model.Article
	s, _ := c.Get("cid")
	_ = c.ShouldBindJSON(&data)
	fmt.Println(s)
	code = model.CreateArt(&data)
	//给前端返回的数据
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//  查询分类下的所有文章
func GetArtilcesByCateId(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	if pageNum == 0 {
		pageNum = 1
	}
	data, code, total := model.GetCateArt(pageSize, pageNum, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//  查询单个文章信息 通过文章id
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code, data := model.GetArtInfo(id)
	//fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询所有文章
func GetAllArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageNum == 0 {
		pageNum = 1
	}
	data, code, total := model.GetArt(pageSize, pageNum)
	//给前端返回的数据
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑文章
func EditArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	fmt.Println(data)
	code = model.EditArt(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

//删除文章
func DelArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("qqqq", id)
	code = model.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

func Write(c *gin.Context) {
	c.HTML(http.StatusOK, "writing.html", gin.H{
		//"CdnURL":    utils.CdnUrl,
		"Title":     "forrest",
		"Categorys": model.GetCategory(),
	})
}

func GetDetail(c *gin.Context) {
	path := c.Request.URL.String()
	IdStr := strings.TrimPrefix(path, "/article/")
	IdStr = strings.TrimSuffix(IdStr, ".html")
	id, _ := strconv.Atoi(IdStr)
	_, data := model.GetArtInfo(id)
	//fmt.Println("QQQQ", c.Request.URL)
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"Title":       utils.Title,
		"Description": utils.Description,
		"Logo":        utils.Logo,
		"Navigation":  utils.Navigation,
		"Title1":      model.GetSingleSpecArt(data).Title,
		"Cid1":        model.GetSingleSpecArt(data).Cid,
		"Name1":       model.GetSingleSpecArt(data).Name,
		"ID1":         model.GetSingleSpecArt(data).ID,
		"CreatedAt1":  model.GetSingleSpecArt(data).CreatedAt,
		"Content1":    template.HTML(model.GetSingleSpecArt(data).Content),
	})
}

func GetArticleByCategoryId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.Query("page"))

	arts, _, _ := model.GetCateArt(10, page, id)
	_, _, total := model.GetCateArt(10000, 1, id)
	artsSpec := model.GetArtSpec(arts)
	Name := artsSpec[0].Name
	pageCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < int(pageCount); i++ {
		pages = append(pages, i+1)
	}
	c.HTML(http.StatusOK, "category.html", gin.H{
		"Title":        utils.Title,
		"Description":  utils.Description,
		"Logo":         utils.Logo,
		"Navigation":   utils.Navigation,
		"Articles":     artsSpec,
		"Total":        total,
		"Page":         page,
		"Pages":        pages,
		"PageEnd":      page != int(pageCount),
		"CategoryName": Name,
	})
}
