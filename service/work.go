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

//GetWorksNoLimit 前端读取所有作品
func GetWorksNoLimit() (works []*model.Work, err error) {

	return db.GetWorksNoLimit()
}

//AddWork 新增作品--事务，多表
func AddWork(userID int, work *model.WorkDetail) (int, error) {
	work.UserID = userID
	return db.AddWork(userID, work)
}

//EditWork 修改作品--事务，多表
func EditWork(userID int, work *model.WorkDetail) (err error) {
	return db.EditWork(userID, work)
}

//RADelWorksFE 永久删除作品--事务，多表
func RADelWorksFE(ids []int) (err error) {
	return db.DelWorksFE(ids)
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

//RACheckWorks 作品审核
func RACheckWorks(ids []int, status int) (err error) {
	return db.CheckWorks(ids, status)
}

//CountWork 访问数量
func CountWork(id int) (err error) {
	return db.CountWork(id)
}

//新增服务

//RAGetWork 获取文章详情
func RAGetWork(workID int) (*model.WorkDetail, error) {
	return db.GetRPWork(workID)
}

//RAGetAllWorksByStatus 通过状态获取所有文章列表
func RAGetAllWorksByStatus(page int, args ...interface{}) ([]*model.Work, error) {
	offset := (page - 1) * setting.PageSize
	return db.GetAllWorksByStatus(offset, setting.PageSize, args...)
}

//TotalWorkByAuthor 根据userID来计算文章
func TotalWorkByAuthor(userID int) (int, error) {
	return db.TotalWorkByAuthor(userID)
}

//TotalWork 计算作品记录，审核通过，或者未删除的数量。
func TotalWork() (int, error) {
	return db.TotalWork()
}

//TotalWorkByTag 计算作品总数
func TotalWorkByTag(tagID int) (num int, err error) {
	return db.TotalWorkByTag(tagID)
}

//RATotalWork 管理员计算
func RATotalWork(args ...interface{}) (int, error) {
	return db.TotalWorkByStatus(args...)
}

//RAExistWork 判断文章id是否真实存在,无论状态
func RAExistWork(workID int) (bool, error) {
	return db.ExistWork(workID)
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
