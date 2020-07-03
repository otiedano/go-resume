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
func DelArticleCategories(articleCategories *model.ArticleCategories) (err error) {

	var ids []int
	for _, v := range *articleCategories {
		if v.CategoryID != 0 {
			ids = append(ids, v.CategoryID)
		}
	}

	return db.DelArticleCategories(ids)
}

//ExistArticleCategory 判断分类是否存在
func ExistArticleCategory(id int) (bool, error) {
	return db.ExistArticleCategory(id)
}
