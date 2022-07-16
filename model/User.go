package model

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
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
	//data.Password = ScryptPassword(data.Password)
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

//使用Gorm的钩子函数进行加密
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPassword(u.Password)
	return
}

//密码加密
func ScryptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{'m', 'a', 'g', 'u', 'a', 'n', 'g', 'y'}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	//放在数据库里的密码 Fpw
	Fpw := base64.StdEncoding.EncodeToString(HashPw)
	return Fpw
}
