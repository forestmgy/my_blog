package model

import (
	"gorm.io/gorm"
	"my_blog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`  //category id 文章对应分类的id
	Desc    string `gorm:"type:varchar(200)" json:"desc"` //Description --文章描述
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

//新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo 查询 文章列表
func GetArt(pageSize, pageNum int) []Category { //pageSize --每页最大数量  pageNum -- 当前页数
	var cates []Category
	db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates)
	return cates
}

//todo 查询分类下的所有文章

//todo 查询单个文章

//编辑文章
func EditArt(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&Article{}).Where("ID = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除文章
func DeleteArt(id int) int {
	var cate Article
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
