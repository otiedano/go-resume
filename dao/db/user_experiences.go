package db

import (
	"database/sql"
	"errors"
	"fmt"
	"sz_resume_202005/model"
	"sz_resume_202005/utils/zlog"

	"github.com/jmoiron/sqlx"
)

//GetExperiences 查看工作经历
func GetExperiences(id int) (userExp []*model.Experience, err error) {
	sqlStr := "select company, content,start_time,end_time,salary,leave_reason,user_id,exp_id from experience where user_id=? order by start_time asc"
	err = db.Select(&userExp, sqlStr, id)
	return
}

//AddExperience 新增单个工作经历
func AddExperience(exp *model.Experience) (err error) {
	sqlStr := "insert into experience (company,content,start_time,end_time,salary,leave_reason,user_id) values (?)"
	s, args, err := sqlx.In(sqlStr, exp)
	if err != nil {
		return err
	}
	result, err := db.Exec(s, args...)

	if err != nil {
		return err
	}
	rs, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rs != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return
}

//AddExperiences 新增工作经历 不用事务，不更新用户表状态
func AddExperiences(exps *model.Experiences) (err error) {

	l := len(*exps)
	if l < 1 {
		return
	}

	strParam := batchStringParam(l)
	sqlStr := "insert into experience (company,content,start_time,end_time,salary,leave_reason,user_id) values " + strParam

	s, args, err := sqlx.In(sqlStr, exps.ConvInterfaceArray()...)
	if err != nil {
		return err
	}
	result, err := db.Exec(s, args...)

	if err != nil {
		return err
	}
	rs, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rs < 1 {
		return errors.New("exec sqlStr1 failed")
	}

	return nil
}

// //AddExpx 新增工作经历，带事务更新has_experience
// func AddExpx(exps *model.Experiences) (err error) {

// 	l := len(*exps)
// 	if l < 1 {
// 		return
// 	}

// 	tx, err := db.Beginx() // 开启事务
// 	if err != nil {
// 		return err
// 	}
// 	defer func() {
// 		//recover()函数必须在defer内执行，其检查程序是否产生panic。
// 		if p := recover(); p != nil {
// 			tx.Rollback()
// 			panic(p) // re-throw panic after Rollback
// 		} else if err != nil {
// 			zlog.Errorf("trans failed, err:%v\n", err)
// 			zlog.Debug("run rollback")
// 			err = tx.Rollback() // err is non-nil; don't change it
// 			if err != nil {
// 				zlog.Fatalf("transaction err:%v", err)
// 			}
// 		}
// 	}()

// 	strParam := batchStringParam(l)
// 	sqlStr := "insert into experience (company,content,start_time,end_time,salary,leave_reason,user_id) values " + strParam

// 	s, args, err := sqlx.In(sqlStr, exps.ConvInterfaceArray()...)
// 	if err != nil {
// 		return err
// 	}
// 	rs, err := tx.Exec(s, args...)

// 	if err != nil {
// 		return err
// 	}
// 	n, err := rs.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if n < 1 {
// 		return errors.New("exec sqlStr1 failed")
// 	}

// 	//安全起见，应该遍历所有的数组，但是不做这么复杂的逻辑，直接默认传入的都是同一用户的。
// 	sqlStr2 := "update user set has_experience=true where user_id=?"

// 	rs, err = tx.Exec(sqlStr2, (*exps)[0].UserID)

// 	if err != nil {
// 		return err
// 	}
// 	n, err = rs.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if n < 1 {
// 		return errors.New("exec sqlStr2 failed")
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//EditExperience 修改工作经历
func EditExperience(exp *model.Experience) (err error) {

	sqlStr := "update experience set company=?,content=?,start_time=?,end_time=?,salary=?,leave_reason=? where user_id=? and exp_id=?"

	result, err := db.Exec(sqlStr, exp.Company, exp.Content, exp.StartTime, exp.EndTime, exp.Salary, exp.LeaveReason, exp.UserID, exp.ExpID)
	if err != nil {
		return
	}
	rs, err := result.RowsAffected()

	if err != nil {
		return
	}
	if rs != 1 {
		return fmt.Errorf("update rowsAffected not equal than 1")
	}
	return nil
}

//DelExperience 删除工作经历
// func DelExperience(id int) (err error) {
// 	sqlStr := "delete from experience where id"
// 	result, err := db.Exec(sqlStr)
// 	if err != nil {
// 		return
// 	}
// 	rs, err := result.RowsAffected()
// 	if err != nil {
// 		return
// 	}
// 	if rs != 1 {
// 		return fmt.Errorf("delete rowsAffected not equal than 1")
// 	}
// 	return nil
// }

//DelExperiences 删除工作经历
func DelExperiences(userID int, ids []int) (err error) {
	sqlStr := "delete from experience where exp_id in (?) and user_id=?"
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

//ExistExperience 判断experience是否存在。
func ExistExperience(exp *model.Experience) (bool, error) {
	sqlStr := "select exp_id from experience where exp_id=? and user_id=? "
	var rexp model.Experience
	err := db.Get(&rexp, sqlStr, exp.ExpID, exp.UserID)
	zlog.Debug("rexp", rexp)
	zlog.Debug("err", err)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, err
}

//GetExperience 查看工作经历
func GetExperience(userID int, expID int) (exp *model.Experience, err error) {
	exp = &model.Experience{}
	sqlStr := "select company, content,start_time,end_time,salary,leave_reason,user_id,exp_id from experience where user_id=? and exp_id=?"
	err = db.Get(exp, sqlStr, userID, expID)
	return
}
