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

		//Article模块的路由接口

	}
	r.Run(utils.HttpPort)
}
