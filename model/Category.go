package model

import "my_blog/utils/errmsg"

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在--通过用户id来查
func CheckCategory(categoryid int) int {
	var cate Category
	db.Select("id").Where("id= ?", categoryid).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_DULPICATED
	}
	return errmsg.SUCCESS
}

//查询分类名是否被用过了
func CheckCategoryName(name string) int {
	var cate Category
	db.Select("id").Where("Name= ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_DULPICATED
	}
	return errmsg.SUCCESS
}

//新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类列表
func GetCategory(pageSize, pageNum int) []Category { //pageSize --每页最大数量  pageNum -- 当前页数
	var cates []Category
	db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates)
	return cates
}

//todo 查询分类下的所有文章

//编辑分类
func EditCategory(id int, data *Category) int {
	code := CheckCategory(id)
	if code == errmsg.SUCCESS {
		return errmsg.ERROR_CATENAME_NOT_EXIST
	}
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&Category{}).Where("ID = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除分类
func DeleteCategory(id int) int {
	code := CheckCategory(id)
	if code == errmsg.SUCCESS {
		return errmsg.ERROR_CATENAME_NOT_EXIST
	}
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
