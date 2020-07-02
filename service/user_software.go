package service

import (
	"sz_resume_202005/dao/db"
	"sz_resume_202005/model"
)

//GetSoftwares 查看软件
func GetSoftwares(userID int) ([]*model.Software, error) {
	return db.GetSoftwares(userID)
}

//AddSoftware 新增软件
func AddSoftware(userID int, software *model.Software) error {
	software.UserID = userID
	return db.AddSoftware(software)
}

//EditSoftware 编辑软件
func EditSoftware(userID int, software *model.Software) error {
	software.UserID = userID
	return db.EditSoftware(software)
}

//ExistSoftware 编辑软件
func ExistSoftware(userID int, software *model.Software) (bool, error) {
	software.UserID = userID

	return db.ExistSoftware(software)
}

//DelSoftwares 删除软件
func DelSoftwares(userID int, ids []int) error {
	return db.DelSoftwares(userID, ids)
}