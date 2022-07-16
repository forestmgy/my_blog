package model

import (
	"gorm.io/gorm"
	"my_blog/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

//查询用户是否存在--通过用户名来查
func CheckUser(username string) int {
	var users User
	db.Select("id").Where("username= ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_DUPLICATED
	}
	return errmsg.SUCCESS
}

//注册用户
func CreateUser(data *User) int {
	err := db.Create(&data)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表并分页
func GetUsers(pageSize, pageNum int) []User { //pageSize --每页最大数量  pageNum -- 当前页数
	var users []User
	db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	return users
}

//编辑用户

//删除用户
