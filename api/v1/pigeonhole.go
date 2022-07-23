package v1

import (
	"github.com/gin-gonic/gin"
	"my_blog/model"
	"my_blog/utils"
	"net/http"
)

func GetPigeonhole(c *gin.Context) {
	pigeonholeMap := make(map[string][]model.ArticleList)
	arts, _ := model.GetAllArt()
	artsSpec := model.GetArtSpec(arts)
	for _, art := range artsSpec {
		at := art.CreatedAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], art)
	}
	c.HTML(http.StatusOK, "pigeonhole.html", gin.H{
		"Title":       utils.Title,
		"Description": utils.Description,
		"Logo":        utils.Logo,
		"Lines":       pigeonholeMap,
		"Navigation":  utils.Navigation,
	})
}
