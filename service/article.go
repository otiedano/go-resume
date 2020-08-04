package service

import (
	"sz_resume_202005/dao/db"
	"sz_resume_202005/model"
	"sz_resume_202005/utils/setting"
)

//GetArticle 读取审核过的文章详情
func GetArticle(articleID int) (article *model.ArticleDetail, err error) {

	return db.GetArticle(articleID)
}

//GetPArticle 读取文章详情
func GetPArticle(userID, articleID int) (article *model.ArticleDetail, err error) {
	return db.GetPArticle(userID, articleID)
}

//GetArticlesByCategory 通过分类和分页获取文章
func GetArticlesByCategory(categoryID int, page int) ([]*model.Article, error) {

	offset := (page - 1) * setting.PageSize
	return db.GetArticlesByCategory(categoryID, offset, setting.PageSize)
}

//GetArticlesByAuthor 通过用户获取文章，后端
func GetArticlesByAuthor(userID int, page int) ([]*model.Article, error) {
	offset := (page - 1) * setting.PageSize
	return db.GetArticlesByAuthor(userID, offset, setting.PageSize)
}

//AddArticle 新增文章
func AddArticle(userID int, article *model.ArticleDetail) (int, error) {
	return db.AddArticle(userID, article)
}

//EditArticle 修改文章
func EditArticle(userID int, article *model.ArticleDetail) (err error) {
	return db.EditArticle(userID, article)
}

//RADelArticlesFE 删除文章
func RADelArticlesFE(ids []int) (err error) {
	return db.DelArticlesFE(ids)
}

//DelArticles 删除文章
func DelArticles(userID int, ids []int) (err error) {
	return db.DelArticles(userID, ids)
}

//ExistArticleByAuth 检查是否有操作文章的权限
func ExistArticleByAuth(userID int, articleID int) (bool, error) {

	return db.ExistArticleByAuth(userID, articleID)
}

//ExistArticleByID 前端判断文章是否存在
func ExistArticleByID(id int) (bool, error) {

	return db.ExistArticleByID(id)
}

//RACheckArticles 文章审核
func RACheckArticles(ids []int, status int) (err error) {
	return db.CheckArticles(ids, status)
}

//CountArticle 访问数量
func CountArticle(id int) (err error) {
	return db.CountArticle(id)
}

//新增服务

//RAGetArticle 获取文章详情
func RAGetArticle(articleID int) (*model.ArticleDetail, error) {
	return db.GetRPArticle(articleID)
}

//RAGetAllArtilesByStatus 通过状态获取所有文章列表
func RAGetAllArtilesByStatus(page int, args ...interface{}) ([]*model.Article, error) {
	offset := (page - 1) * setting.PageSize
	return db.GetAllArtilesByStatus(offset, setting.PageSize, args...)
}

//TotalArticleByAuthor 根据userID来计算文章
func TotalArticleByAuthor(userID int) (int, error) {
	return db.TotalArticleByAuthor(userID)
}

//TotalArticle 获取审核通过的内容
func TotalArticle() (int, error) {
	return db.TotalArticle()
}

//TotalArticleByCategory 获取审核通过的内容
func TotalArticleByCategory(categoryID int) (int, error) {
	return db.TotalArticleByCategory(categoryID)
}

//RATotalArticle 管理员计算
func RATotalArticle(args ...interface{}) (int, error) {
	return db.TotalArticleByStatus(args...)
}

//RAExistArticle 判断文章id是否真实存在,无论状态
func RAExistArticle(articleID int) (bool, error) {
	return db.ExistArticle(articleID)
}

// //SetArticleCount 记录访问数量，同一ip只能访问一次
// func SetArticleCount(origin, pid string, id int) (err error) {

// 	name := fmt.Sprintf("articles_%s_%s", pid, origin)
// 	_, exist, err := redisdb.Get(name)
// 	if err != nil {
// 		zlog.Error(err)

// 		return
// 	}
// 	if !exist {
// 		CountArticle(id)
// 		err = redisdb.Set(name, "1", 600)
// 		return
// 	}
// 	tm, err := redisdb.GetTTL(name)
// 	if err != nil {
// 		zlog.Error(err)
// 		return
// 	}
// 	if tm < 0 {
// 		CountArticle(id)
// 		err = redisdb.Set(name, "1", 600)
// 	}
// 	return
// }
