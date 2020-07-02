package db

import (
	"fmt"
	"sz_resume_202005/model"
	"time"

	"github.com/jmoiron/sqlx"
)

//标签管理部分

//AddTag 新增标签
func AddTag(tag *model.WorkTag) (err error) {
	sqlStr := `
	INSERT INTO work_tag (tag_name,tag_no) VALUES (?,?)
   `
	_, err = db.Exec(sqlStr, tag.TagName, tag.TagNO)
	return
}

//EditTag 编辑标签
func EditTag(tag *model.WorkTag) (err error) {
	sqlStr := `
	UPDATE work_tag SET tag_name=?,tag_no=?,update_time=? WHERE tag_id=?
   `
	_, err = db.Exec(sqlStr, tag.TagName, tag.TagNO, time.Now(), tag.TagID)
	return
}

//EditTags 编辑标签组
func EditTags(tags *model.WorkTags) (err error) {
	l := len(*tags)

	if l < 1 {
		return
	}

	sqlStr := fmt.Sprintf(`
	INSERT INTO work_tag (tag_id,tag_name,tag_no) VALUES %s 
	ON DUPLICATE KEY 
	UPDATE tag_name=VALUES(tag_name),tag_no=VALUES(tag_no)`, batchStringParam(l))

	query, args, err := sqlx.In(sqlStr, tags.ConvInterfaceArray()...)
	query += ",update_time=?"
	args = append(args, time.Now())

	if err != nil {
		return
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return
	}
	rs, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rs < 1 {
		err = fmt.Errorf("rs not equal 1")
		return
	}
	return
}

//DelTags 删除标签
func DelTags(ids []int) (err error) {
	sqlStr := `
   DELETE FROM work_tag WHERE tag_id in (?)
  `
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return
	}
	_, err = db.Exec(query, args...)
	return

}

//GetTags 获取标签
func GetTags() (tags []*model.WorkTag, err error) {
	sqlStr := `
  SELECT tag_id,tag_name,tag_no,create_time,update_time FROM work_tag  ORDER BY tag_no desc
  `
	err = db.Select(&tags, sqlStr)
	return
}
