package main

import (
	"my_blog/model"
	"my_blog/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	routes.InitRouter()
}
