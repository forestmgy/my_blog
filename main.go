package main

import (
	"fmt"
	"my_blog/model"
	"my_blog/routes"
	"my_blog/utils"
)

func main() {
	//引用数据库
	fmt.Println(utils.CdnUrl)
	model.InitDb()
	routes.InitRouter()
}
