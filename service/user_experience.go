package service

import (
	"sz_resume_202005/dao/db"
	"sz_resume_202005/model"
)

//GetExperiences 查看用户工作经历
func GetExperiences(userID int) ([]*model.Experience, error) {
	return db.GetExperiences(userID)

}

//GetExperience 查看工作经历
func GetExperience(userID int, expID int) (exp *model.Experience, err error) {
	return db.GetExperience(userID, expID)
}

//AddExperiences 新增工作经历
func AddExperiences(userID int, exps model.Experiences) error {
	for _, v := range exps {
		v.UserID = userID
	}
	return db.AddExperiences(&exps)

}

//AddExperience 新增工作经历
func AddExperience(userID int, exp *model.Experience) error {
	exp.UserID = userID
	return db.AddExperience(exp)

}

//EditExperience 编辑工作经验
func EditExperience(userID int, exp *model.Experience) error {
	exp.UserID = userID

	return db.EditExperience(exp)
}

//ExistExperience 编辑工作经验
func ExistExperience(userID int, exp *model.Experience) (bool, error) {
	exp.UserID = userID

	return db.ExistExperience(exp)
}

// //DelExperience 删除工作经验
// func DelExperience(id int) error {
// 	return db.DelExperience(id)

// }

//DelExperiences 删除工作经验
func DelExperiences(userID int, ids []int) error {
	return db.DelExperiences(userID, ids)
}
