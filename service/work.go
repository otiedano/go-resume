package service

import (
	"sz_resume_202005/dao/db"
	"sz_resume_202005/model"
	"sz_resume_202005/utils/setting"
)

//GetWork 读取审核过的作品详情
func GetWork(workID int) (w *model.WorkDetail, err error) {
	return db.GetWork(workID)
}

//GetPWork 读取所有状态的作品详情
func GetPWork(userID, workID int) (w *model.WorkDetail, err error) {
	return db.GetPWork(userID, workID)
}

//GetWorksByAuthor 根据作者返回作品列表,用于后端管理
func GetWorksByAuthor(userID int, page int) (works []*model.Work, err error) {
	offset := (page - 1) * setting.PageSize
	return db.GetWorksByAuthor(userID, offset, setting.PageSize)
}

//GetWorksByTag 根据分页和分类返回作品列表，用于前端显示
func GetWorksByTag(tagID int, page int) (works []*model.Work, err error) {
	offset := (page - 1) * setting.PageSize
	return db.GetWorksByTag(tagID, offset, setting.PageSize)
}

//AddWork 新增作品--事务，多表
func AddWork(userID int, work *model.WorkDetail) (err error) {
	work.UserID = userID
	return db.AddWork(userID, work)
}

//EditWork 修改作品--事务，多表
func EditWork(userID int, work *model.WorkDetail) (err error) {
	return db.EditWork(userID, work)
}

//DelWorksFE 永久删除作品--事务，多表
func DelWorksFE(userID int, ids []int) (err error) {
	return db.DelWorksFE(userID, ids)
}

//DelWorks 删除作品--软删除 --事务，多表
func DelWorks(userID int, ids []int) (err error) {
	return db.DelWorks(userID, ids)
}

//ExistWorkByAuth 检查是否有操作作品的权限
func ExistWorkByAuth(userID, workID int) (bool, error) {

	return db.ExistWorkByAuth(userID, workID)
}

//ExistWorkByID 前端判断作品是否存在
func ExistWorkByID(workID int) (bool, error) {

	return db.ExistWorkByID(workID)
}

//CheckWorks 作品审核
func CheckWorks(userID int, ids []int, status int) (err error) {
	return db.CheckWorks(userID, ids, status)
}

//CountWork 访问数量
func CountWork(id int) (err error) {
	return db.CountWork(id)
}

//新增服务

//GetRPWork 获取文章详情
func GetRPWork(userID, articleID int) (*model.ArticleDetail, error) {
	return db.GetRPArticle(userID, articleID)
}

//GetAllWorksByStatus 通过状态获取所有文章列表
func GetAllWorksByStatus(userID int, page int, args ...interface{}) ([]*model.Article, error) {
	offset := (page - 1) * setting.PageSize
	return db.GetAllArtilesByStatus(userID, offset, setting.PageSize, args...)
}

//TotalWorkByAuthor 根据userID来计算文章
func TotalWorkByAuthor(userID int, args ...interface{}) {
	db.TotalArticleByAuthor(userID, args...)
}

//TotalWorkByStatus 管理员计算
func TotalWorkByStatus(userID int, args ...interface{}) {
	db.TotalWorkByStatus(userID, args...)
}

//ExistWork 判断文章id是否真实存在,无论状态
func ExistWork(userID int, workID int) (bool, error) {
	return db.ExistWork(userID, workID)
}

// //SetWorkCount (origin,)
// func SetWorkCount(origin, pid string, id int) (err error) {

// 	name := fmt.Sprintf("works_%s_%s", pid, origin)
// 	_, exist, err := redisdb.Get(name)
// 	if err != nil {
// 		zlog.Error(err)

// 		return
// 	}
// 	if !exist {
// 		CountWork(id)
// 		err = redisdb.Set(name, "1", 600)
// 		return
// 	}
// 	tm, err := redisdb.GetTTL(name)
// 	if err != nil {
// 		zlog.Error(err)
// 		return
// 	}
// 	if tm < 0 {
// 		CountWork(id)
// 		err = redisdb.Set(name, "1", 600)
// 	}
// 	return
// }
