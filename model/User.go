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

//查询用户是否存在--通过用户id来查
func CheckUser(userid int) int {
	var users User
	db.Select("id").Where("id= ?", userid).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_DUPLICATED
	}
	return errmsg.SUCCESS
}

//查询用户名是否被用过了
func CheckUserName(username string) int {
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
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表并分页
func GetUsers(pageSize, pageNum int) ([]User, int64) { //pageSize --每页最大数量  pageNum -- 当前页数
	var users []User
	var total int64
	db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total)
	return users, total
}

//编辑用户
func EditUser(id int, data *User) int {
	code := CheckUser(id)
	if code == errmsg.SUCCESS {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
func DeleteUser(id int) int {
	code := CheckUser(id)
	if code == errmsg.SUCCESS {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

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

//登录验证
func CheckLogin(username, password string) int {
	var user User

	db.Where("username =?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	if ScryptPassword(password) != user.Password {
		return errmsg.ERROE_USERPASSWORD_WRONG
	}
	return errmsg.SUCCESS
}
