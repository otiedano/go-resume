package db

import (
	"fmt"
	"sz_resume_202005/model"
	"testing"
)

//标签管理部分

//AddTag 新增标签
func TestAddTag(t *testing.T) {
	tag := &model.WorkTag{
		TagName: "Flash",
	}
	_, err := AddTag(tag)
	if err != nil {
		t.Error("err:", err)
	}
}

//EditTag 编辑标签
func TestEditTag(t *testing.T) {
	tag := &model.WorkTag{
		TagName: "UI/前端",
		TagID:   1,
		TagNO:   4,
	}
	err := EditTag(tag)
	if err != nil {
		t.Error("err:", err)
	}
}

func TestEditTags(t *testing.T) {
	tags := model.WorkTags{}
	tags = append(tags, &model.WorkTag{
		TagName: "U端",
		TagID:   1,
		TagNO:   4,
	}, &model.WorkTag{
		TagName: "后面",
		TagID:   2,
		TagNO:   3,
	})
	err := EditTags(&tags)
	if err != nil {
		t.Error("err:", err)
	}
}

//DelTags 删除标签
func TestDelTags(t *testing.T) {
	ids := []int{1, 2}
	err := DelTags(ids)
	if err != nil {
		t.Error("err", err)
	}
}

//GetTags 获取标签
func TestGetTags(t *testing.T) {
	tags, err := GetTags()
	if err != nil {
		t.Error("err:", err)
	}
	for _, v := range tags {
		fmt.Printf("tag:%+v\n", v)
	}
}
func TestExistWorkTag(t *testing.T) {

	b, err := ExistWorkTag(5)
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("not expect result")
	}
}
