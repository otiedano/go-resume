package service

import (
	"sz_resume_202005/dao/db"
	"sz_resume_202005/model"
)

//GetSkills 查看技能
func GetSkills(userID int) ([]*model.Skill, error) {
	return db.GetSkills(userID)
}

//AddSkill 新增技能
func AddSkill(userID int, skill *model.Skill) error {
	skill.UserID = userID
	return db.AddSkill(skill)
}

//EditSkill 编辑技能
func EditSkill(userID int, skill *model.Skill) error {
	skill.UserID = userID
	return db.EditSkill(skill)
}

//ExistSkill 技能是否存在
func ExistSkill(userID int, skill *model.Skill) (bool, error) {
	skill.UserID = userID

	return db.ExistSkill(skill)
}

//DelSkills 删除技能
func DelSkills(userID int, ids []int) error {
	return db.DelSkills(userID, ids)
}
