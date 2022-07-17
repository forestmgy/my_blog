package routes

import (
	"github.com/gin-gonic/gin"
	v1 "my_blog/api/v1"
	"my_blog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	routerv1 := r.Group("api/v1")
	{
		//User模块的路由接口
		routerv1.POST("user/add", v1.AddUser)
		routerv1.GET("users", v1.GetAllUser)
		routerv1.PUT("user/:id", v1.EditUser)
		routerv1.DELETE("user/:id", v1.DelUser)
		//Category模块的路由接口
		routerv1.POST("category/add", v1.AddCategory)
		routerv1.GET("categorys", v1.GetAllCategory)
		routerv1.PUT("category/:id", v1.EditCategory)
		routerv1.DELETE("category/:id", v1.DelCategory)
		//Article模块的路由接口
		routerv1.POST("article/add", v1.AddArt)                      //添加文章
		routerv1.GET("articles", v1.GetAllArt)                       //得到全部文章
		routerv1.PUT("article/:id", v1.EditArt)                      //编辑文章
		routerv1.DELETE("article/:id", v1.DelArt)                    //删除文章
		routerv1.GET("article/:id", v1.GetArtInfo)                   //得到单个文章
		routerv1.GET("article/catelist/:id", v1.GetArtilcesByCateId) //通过分类名ID得到文章列表
	}
	r.Run(utils.HttpPort)
}
