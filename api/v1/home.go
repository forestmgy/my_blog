package v1

import (
	"github.com/gin-gonic/gin"
	"my_blog/model"
	"my_blog/utils"
	"net/http"
	"strconv"
)

func GetHome(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	arts, _, _ := model.GetArt(10, page)
	_, _, total := model.GetArt(100, 1)
	artsSpec := model.GetArtSpec(arts)
	pageCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < int(pageCount); i++ {
		pages = append(pages, i+1)
	}

	//fmt.Println(page, "qqqq", total, pages)
	c.HTML(http.StatusOK, "template/index.html", gin.H{
		"Title":       utils.Title,
		"Description": utils.Description,
		"Logo":        utils.Logo,
		"Github":      utils.Github,
		"Avatar":      utils.Avatar,
		"UserName":    utils.UserName,
		"UserDesc":    utils.UserDesc,
		"Categorys":   model.GetCategory(),
		"Navigation":  utils.Navigation,
		"Articles":    artsSpec,
		"Total":       total,
		"Page":        page,
		"Pages":       pages,
		"PageEnd":     page != int(pageCount),
	})

}
