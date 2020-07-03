package service

import (
	"sz_resume_202005/dao/db"
	"sz_resume_202005/model"
)

//AddWorkTag 新增标签
func AddWorkTag(tag *model.WorkTag) (int, error) {
	return db.AddTag(tag)
}

//EditWorkTag 编辑标签
func EditWorkTag(tag *model.WorkTag) (err error) {
	return db.EditTag(tag)
}

//EditWorkTags 编辑标签组
func EditWorkTags(tags *model.WorkTags) (err error) {

	return db.EditTags(tags)
}

//DelWorkTags 删除标签
func DelWorkTags(workTags *model.WorkTags) (err error) {

	var ids []int

	for _, v := range *workTags {
		if v.TagID != 0 {
			ids = append(ids, v.TagID)
		}
	}
	return db.DelTags(ids)
}

//GetWorkTags 获取标签
func GetWorkTags() (tags []*model.WorkTag, err error) {
	return db.GetTags()
}

//ExistWorkTag 判断分类是否存在
func ExistWorkTag(id int) (bool, error) {
	return db.ExistWorkTag(id)
}
