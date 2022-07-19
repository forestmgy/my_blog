package utils

import (
	"gopkg.in/ini.v1"
	"log"
	"strings"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string

	Title       string
	Description string
	Logo        string
	Navigation  []string
	Github      string
	Wechat      string
	Avatar      string
	UserName    string
	UserDesc    string
)

func init() {
	file, err := ini.Load("config/config.ini") //取到file文件
	if err != nil {
		log.Println("配置文件读取错误", err)
	}
	//file是一个结构体
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
	LoadInfo(file)
}

func LoadServer(file *ini.File) {
	//从config.ini里取值，取不到就设为默认值“debug
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("m2a0g0u2a1n1g0y5u")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("AppMode").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassword").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("myblog")
}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()
}

func LoadInfo(file *ini.File) {
	Title = file.Section("info").Key("Title").String()
	Description = file.Section("info").Key("Description").String()
	Logo = file.Section("info").Key("Logo").String()
	Navigation = strings.Split(file.Section("info").Key("Navigation").String(), ",")
	Github = file.Section("info").Key("Github").String()
	Wechat = file.Section("info").Key("Wechat").String()
	Avatar = file.Section("info").Key("Avatar").String()
	UserName = file.Section("info").Key("UserName").String()
	UserDesc = file.Section("info").Key("UserDesc").String()
}
