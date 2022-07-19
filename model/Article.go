package model

import (
	"gorm.io/gorm"
	"my_blog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title    string `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int    `gorm:"type:int;not null" json:"cid"`  //category id 文章对应分类的id
	Desc     string `gorm:"type:varchar(200)" json:"desc"` //Description --文章描述
	Content  string `gorm:"type:longtext" json:"content"`
	Markdown string `grom:"type:longtext" json:"markdown"`
}

//新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询 文章列表
func GetArt(pageSize, pageNum int) ([]Article, int, int64) { //pageSize --每页最大数量  pageNum -- 当前页数
	var articleList []Article
	var total int64
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// 查询分类下的所有文章
func GetCateArt(pageSize, pageNum, cid int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", cid).Find(&cateArtList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATENAME_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// 查询单个文章
func GetArtInfo(id int) (int, Article) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return errmsg.ERROR_ART_NOT_EXIST, art
	}
	return errmsg.SUCCESS, art
}

//编辑文章
func EditArt(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
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
