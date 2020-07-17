package db

import (
	"database/sql"
	"errors"
	"fmt"
	"sz_resume_202005/model"
	"sz_resume_202005/utils/zlog"
	"time"

	"github.com/jmoiron/sqlx"
)

//GetSkills 获取skill列表
func GetSkills(userID int) (skills []*model.Skill, err error) {
	sqlStr := "select skill_name,img,per,user_id,skill_id ,skill_no from skill where user_id=? order by skill_no asc"
	err = db.Select(&skills, sqlStr, userID)
	return
}

//GetSkill 获取skill
func GetSkill(userID int, id int) (skill *model.Skill, err error) {

	skill = &model.Skill{}
	sqlStr := "select skill_name,img,per,user_id,skill_id ,skill_no from skill where user_id=? and skill_id=?"

	err = db.Get(skill, sqlStr, userID, id)

	return
}

//AddSkill 单个加skill
func AddSkill(skill *model.Skill) (err error) {
	sqlStr := "insert into skill (skill_name,img,per,user_id,skill_no) values (?,?,?,?,?)"

	result, err := db.Exec(sqlStr, skill.SkillName, skill.Img, skill.Per, skill.UserID, skill.SkillNO)

	if err != nil {
		return err
	}
	rs, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rs != 1 {
		return errors.New("rs not equal 1")
	}
	return
}

//EditSkill 编辑skill
func EditSkill(skill *model.Skill) (err error) {
	sqlStr := "update skill set skill_name=?,img=?,per=?,skill_no=?,update_time=? where user_id=? and skill_id=?"
	st, err := db.Preparex(sqlStr)
	if err != nil {
		return
	}
	defer st.Close()
	result, err := st.Exec(skill.SkillName, skill.Img, skill.Per, skill.SkillNO, time.Now(), skill.UserID, skill.SkillID)
	if err != nil {
		return
	}
	rs, err := result.RowsAffected()
	zlog.Debug("rs:%v", rs)
	if err != nil {
		return
	}
	if rs != 1 {
		return fmt.Errorf("update rowsAffected not equal than 1")
	}
	return nil
}

//EditSkills 批量修改软件
func EditSkills(skills *model.Skills) (err error) {
	l := len(*skills)
	fmt.Printf("skills,%+v\n", skills)
	fmt.Print("l:", l)
	if l < 1 {
		return
	}

	sqlStr := fmt.Sprintf(`
	INSERT INTO skill (skill_id,skill_name,img,per,user_id,skill_no) VALUES %s 
	ON DUPLICATE KEY 
	UPDATE skill_name=VALUES(skill_name),skill_no=VALUES(skill_no),img=VALUES(img),per=VALUES(per)`, batchStringParam(l))
	fmt.Printf("sql:%v\n", sqlStr)
	query, args, err := sqlx.In(sqlStr, skills.ConvInterfaceArray()...)
	query += ",update_time=?"
	args = append(args, time.Now())

	fmt.Printf("sql:%v\n", query)
	fmt.Printf("args:%+v\n", args)
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

//DelSkills 删除多个skill
func DelSkills(userID int, ids []int) (err error) {
	sqlStr := "delete from skill where skill_id in (?) and user_id=?"
	query, args, err := sqlx.In(sqlStr, ids, userID)
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
		return fmt.Errorf("delete rowsAffected not larger than 1")
	}
	return nil
}

//ExistSkill 检查skill是否存在
func ExistSkill(skill *model.Skill) (bool, error) {
	sqlStr := "select skill_id from skill where skill_id=? and user_id=? "
	var rskill model.Skill
	err := db.Get(&rskill, sqlStr, skill.SkillID, skill.UserID)
	zlog.Debug("rskill", rskill)
	zlog.Debug("err", err)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}
