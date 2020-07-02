package db

import (
	"fmt"
	"sz_resume_202005/model"
	"testing"
)

func TestAddSkill(t *testing.T) {
	skill := &model.Skill{

		SkillName: "PHP",
		Img:       "/runtime/upload/images/avatar.jpg",
		Per:       70,
		UserID:    1,
	}
	err := AddSkill(skill)
	if err != nil {
		t.Error("err:", err)
	}

}
func TestGetSkills(t *testing.T) {
	skills, err := GetSkills(1)
	if err != nil {
		t.Error("err:", err)
	}
	for _, v := range skills {
		fmt.Printf("v:%+v\n", v)
	}
}

func TestEditSkill(t *testing.T) {
	skill := &model.Skill{
		SkillName: "newSKill",
		SkillID:   6,
		Img:       "/runtion/upload/images/avatar.jpg",
		Per:       66,
		UserID:    1,
	}
	err := EditSkill(skill)
	if err != nil {
		t.Error("err:", err)
	}
}
func TestEditSkills(t *testing.T) {
	skills := make(model.Skills, 0)
	skills = append(skills, &model.Skill{
		SkillID:   19,
		SkillName: "PS",
		Img:       "/runtime/upload/images/1.jpg",
		Per:       65,
		UserID:    1,
	}, &model.Skill{
		SkillID:   21,
		SkillName: "DW",
		SkillNO:   20,
		Img:       "runtime/haha",
		Per:       55,
		UserID:    1,
	})
	EditSkills(&skills)
}
func TestDelSkills(t *testing.T) {
	ids := []int{8, 9}
	err := DelSkills(1, ids)
	if err != nil {
		t.Error("err:", err)
	}
}

func TestExistSkill(t *testing.T) {
	skill := &model.Skill{
		SkillID: 7,
		UserID:  1,
	}
	b, err := ExistSkill(skill)
	if err != nil {
		t.Error("err:", err)
	}
	if !b {
		t.Error("result wrong")
	}
}
