package service

import (
	"sz_resume_202005/dao/db"
	"sz_resume_202005/model"
)

//AddArticleCategory 新增文章分类
func AddArticleCategory(c *model.ArticleCategory) (intNum int, err error) {
	return db.AddArticleCategory(c)
}

//GetArticleCategories 读取文章分类
func GetArticleCategories() (categories []*model.ArticleCategory, err error) {
	return db.GetArticleCategories()
}

//EditArticleCategory 修改文章分类
func EditArticleCategory(category *model.ArticleCategory) (err error) {
	return db.EditArticleCategory(category)
}

//EditArticleCategories 批量修改文章分类
func EditArticleCategories(categories *model.ArticleCategories) (err error) {
	return db.EditArticleCategories(categories)
}

//DelArticleCategories 删除文章分类
func DelArticleCategories(ids []int) (err error) {
	return db.DelArticleCategories(ids)
}
