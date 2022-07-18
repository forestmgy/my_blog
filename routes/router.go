package routes

import (
	"github.com/gin-gonic/gin"
	v1 "my_blog/api/v1"
	"my_blog/middleware"
	"my_blog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken()) //需要鉴权的功能
	{
		//User模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DelUser)
		//Category模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DelCategory)
		//Article模块的路由接口
		auth.POST("article/add", v1.AddArt) //添加文章
		//得到全部文章
		auth.PUT("article/:id", v1.EditArt)   //编辑文章
		auth.DELETE("article/:id", v1.DelArt) //删除文章
		//上传文件
		auth.POST("upload", v1.Upload)
	}
	routerv1_pub := r.Group("api/v1") //不需要鉴权---游客也可以做到的
	{
		routerv1_pub.POST("user/add", v1.AddUser)
		routerv1_pub.GET("users", v1.GetAllUser)
		routerv1_pub.GET("categorys", v1.GetAllCategory)
		routerv1_pub.GET("articles", v1.GetAllArt)
		routerv1_pub.GET("article/:id", v1.GetArtInfo)                   //得到单个文章
		routerv1_pub.GET("article/catelist/:id", v1.GetArtilcesByCateId) //通过分类名ID得到文章列表
		routerv1_pub.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
