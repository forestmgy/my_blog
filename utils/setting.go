package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini") //取到file文件
	if err != nil {
		log.Println("配置文件读取错误", err)
	}
	//file是一个结构体
	LoadServer(file)
	LoadServer(file)
}

func LoadServer(file *ini.File) {
	//从config.ini里取值，取不到就设为默认值“debug
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("AppMode").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassword").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("myblog")

}
