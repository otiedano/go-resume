package db

import (
	"fmt"
	"sz_resume_202005/model"
	"testing"
	"time"
)

//GetExperience 查看工作经历
func TestGetExperiences(t *testing.T) {
	exps, err := GetExperiences(1)
	if err != nil {
		t.Error("err:", err)
	}
	for _, v := range exps {
		fmt.Printf("exp:%+v\n", *v)
	}

}

//AddExperience 新增单个工作经历
func TestAddExperience(t *testing.T) {
	st, _ := time.Parse("2006-01-02 15:04:05", "2015-12-27 18:44:55")
	et, _ := time.Parse("2006-01-02 15:04:05", "2016-12-27 18:44:55")
	exp := &model.Experience{
		StartTime:   st,
		EndTime:     et,
		Company:     "实习工作",
		Content:     "天天就是上课，跑腿",
		Salary:      0,
		LeaveReason: "不给钱啊",
		UserID:      1,
	}
	err := AddExperience(exp)
	if err != nil {
		t.Error("err:", err)
	}
}

//AddExperiences
func TestAddExperiences(t *testing.T) {
	var exps model.Experiences
	st1, _ := time.Parse("2006-01-02 15:04:05", "2015-12-27 18:44:55")
	st2, _ := time.Parse("2006-01-02 15:04:05", "2016-12-27 18:44:55")
	st3, _ := time.Parse("2006-01-02 15:04:05", "2017-12-27 18:44:55")
	et1, _ := time.Parse("2006-01-02 15:04:05", "2016-12-27 18:44:55")
	et2, _ := time.Parse("2006-01-02 15:04:05", "2017-12-27 18:44:55")
	et3, _ := time.Parse("2006-01-02 15:04:05", "2018-12-27 18:44:55")
	exps = append(exps, &model.Experience{
		StartTime:   st1,
		EndTime:     et1,
		Company:     "第一家公司",
		Content:     "就是扫扫地倒到水",
		Salary:      200,
		LeaveReason: "换你你不走啊",
		UserID:      1,
	}, &model.Experience{
		StartTime:   st2,
		EndTime:     et2,
		Company:     "第二家公司",
		Content:     "给老板跑腿,替老板挡酒",
		Salary:      500,
		LeaveReason: "喝废了",
		UserID:      1,
	}, &model.Experience{
		StartTime:   st3,
		EndTime:     et3,
		Company:     "第三家公司",
		Content:     "开老板车，泡老板妞",
		Salary:      1000,
		LeaveReason: "那是老板老婆",
		UserID:      1,
	})

	err := AddExperiences(&exps)
	if err != nil {
		t.Error("err:", err)
	}
}

//AddExpx
func TestAddExpx(t *testing.T) {

}

//EditExperience
func TestEditExperience(t *testing.T) {
	st, _ := time.Parse("2006-01-02 15:04:05", "2015-12-27 18:44:55")
	et, _ := time.Parse("2006-01-02 15:04:05", "2016-12-27 18:44:55")
	exp := &model.Experience{
		Salary:    100,
		Content:   "老板不让我乱说",
		ExpID:     38,
		UserID:    1,
		StartTime: st,
		EndTime:   et,
	}
	err := EditExperience(exp)
	if err != nil {
		t.Error(err)
	}
}

//DelExperiences
func TestDelExperiences(t *testing.T) {
	exps, err := GetExperiences(1)
	if err != nil {
		t.Error(err)
	}
	ids := make([]int, 0)
	for _, v := range exps {
		ids = append(ids, v.ExpID)
	}
	err = DelExperiences(1, ids)
	if err != nil {
		t.Error("DelExperiences function err:", err)
	}
}

//ExistExperience
func TestExistExperience(t *testing.T) {
	exp := &model.Experience{
		UserID: 1,
		ExpID:  38,
	}
	b, err := ExistExperience(exp)
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("b is not expect result")
	}
}
