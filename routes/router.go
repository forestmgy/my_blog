package routes

import (
	"github.com/gin-gonic/gin"
	"html/template"
	v1 "my_blog/api/v1"
	"my_blog/middleware"
	"my_blog/utils"
	"time"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	//r.Use(middleware.Logger())
	r.Static("/resource", "./public/resource")
	r.SetFuncMap(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay, "date1": Date1})
	r.LoadHTMLGlob("template/**/*")
	//r.LoadHTMLFiles("template")
	//r.LoadHTMLGlob("template/pages/*")

	r.GET("/", v1.GetHome)
	r.GET("/login", v1.LoginView)
	r.GET("/writing", v1.Write)
	r.GET("/article/:id.html", v1.GetDetail)
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken()) //需要鉴权的功能
	{
		//User模块的路由接口

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
		//routerv1_pub.GET("categorys", v1.GetAllCategory)
		routerv1_pub.GET("articles", v1.GetAllArt)
		routerv1_pub.GET("article/:id", v1.GetArtInfo)                   //得到单个文章
		routerv1_pub.GET("article/catelist/:id", v1.GetArtilcesByCateId) //通过分类名ID得到文章列表
		routerv1_pub.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
func Date1(date time.Time) string {
	return date.Format("2006-01-02")
}
