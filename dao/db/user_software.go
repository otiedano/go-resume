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

//GetSoftwares 获取softwares
func GetSoftwares(userID int) (softwares []*model.Software, err error) {
	sqlStr := "select software_name,img,user_id,software_id,software_no from software where user_id=? order by software_no asc"
	err = db.Select(&softwares, sqlStr, userID)
	return
}

//GetSoftware 获取skill
func GetSoftware(userID int, id int) (software *model.Software, err error) {
	software = &model.Software{}
	sqlStr := "select software_name,img,user_id,software_id,software_no from software where user_id=? and software_id=?"
	err = db.Select(software, sqlStr, userID, id)
	return
}

//AddSoftware 单个加software
func AddSoftware(software *model.Software) (err error) {
	sqlStr := "insert into software (software_name,img,user_id,software_no) values (?,?,?,?)"

	result, err := db.Exec(sqlStr, software.SoftwareName, software.Img, software.UserID, software.SoftwareNO)

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

//EditSoftware 编辑software
func EditSoftware(software *model.Software) (err error) {
	sqlStr := "update software set software_name=?,img=?,software_no=?,update_time=? where user_id=? and software_id=?"

	result, err := db.Exec(sqlStr, software.SoftwareName, software.Img, software.SoftwareNO, time.Now(), software.UserID, software.SoftwareID)
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

//EditSoftwares 批量修改软件
func EditSoftwares(softwares *model.Softwares) (err error) {
	l := len(*softwares)
	fmt.Printf("softwares,%+v\n", softwares)
	fmt.Print("l:", l)
	if l < 1 {
		return
	}

	sqlStr := fmt.Sprintf(`
	INSERT INTO software (software_id,software_name,img,user_id,software_no) VALUES %s 
	ON DUPLICATE KEY 
	UPDATE software_name=VALUES(software_name),software_no=VALUES(software_no),img=VALUES(img)`, batchStringParam(l))
	fmt.Printf("sql:%v\n", sqlStr)
	query, args, err := sqlx.In(sqlStr, softwares.ConvInterfaceArray()...)
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

//DelSoftwares 删除多个software
func DelSoftwares(userID int, ids []int) (err error) {
	sqlStr := "delete from software where software_id in (?) and user_id=?"
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

//ExistSoftware 检查software是否存在
func ExistSoftware(software *model.Software) (bool, error) {
	sqlStr := "select software_id from software where software_id=? and user_id=? "
	var rsoftware model.Software
	err := db.Get(&rsoftware, sqlStr, software.SoftwareID, software.UserID)
	zlog.Debug("rsoftware", rsoftware)
	zlog.Debug("err", err)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}
