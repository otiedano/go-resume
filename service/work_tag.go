package service

import (
	"sz_resume_202005/dao/db"
	"sz_resume_202005/model"
)

//AddTag 新增标签
func AddTag(tag *model.WorkTag) (err error) {
	return db.AddTag(tag)
}

//EditTag 编辑标签
func EditTag(tag *model.WorkTag) (err error) {
	return db.EditTag(tag)
}

//EditTags 编辑标签组
func EditTags(tags *model.WorkTags) (err error) {

	return db.EditTags(tags)
}

//DelTags 删除标签
func DelTags(ids []int) (err error) {
	return db.DelTags(ids)
}

//GetTags 获取标签
func GetTags() (tags []*model.WorkTag, err error) {
	return db.GetTags()
}
