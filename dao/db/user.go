package db

import (
	"database/sql"
	"sz_resume_202005/model"
	"sz_resume_202005/utils"
	"sz_resume_202005/utils/zlog"
	"time"
)

//CheckUserAuth 查看是否能登录
func CheckUserAuth(u *model.UserAuth) (ifExist bool, err error) {
	//	sqlStr := "select user_id, user_name, avatar,create_time from user where user_name=? and password=?"
	sqlStr := "select user_id from user where user_name=? and password=?"
	var user model.User
	err = db.Get(&user, sqlStr, u.UserName, utils.EncodeSHA256(u.Password))
	if err != nil && err != sql.ErrNoRows {
		return false, err

	}
	if user.UserID > 0 {
		return true, nil
	}
	return false, nil
}

//CheckUser 查看用户是否存在
// func CheckUser(userName string) (ifExist bool, err error) {
// 	//	sqlStr := "select user_id, user_name, avatar,create_time from user where user_name=? and password=?"
// 	sqlStr := "select user_id from user where user_name=?"
// 	var user model.User
// 	err = db.Get(&user, sqlStr, userName)
// 	if err != nil && err != sql.ErrNoRows {
// 		return false, err

// 	}
// 	if user.UserID > 0 {
// 		return true, nil
// 	}
// 	return false, nil
// }

//GetUser 获取用户基本信息
func GetUser(u *model.UserAuth) (*model.User, error) {
	sqlStr := "select user_id, user_name, avatar,create_time from user where user_name=? and password=?"
	//sqlStr := "select user_id from user where user_name=?"
	var user model.User
	err := db.Get(&user, sqlStr, u.UserName, utils.EncodeSHA256(u.Password))
	return &user, err
}

//AddUser 新增用户
func AddUser(userName string, password string) (theID int64, err error) {

	sqlStr := "insert into user(user_name, password,introduce) values (?,?,'')"
	var ret sql.Result
	ret, err = db.Exec(sqlStr, userName, utils.EncodeSHA256(password))
	if err != nil {
		zlog.Errorf("insert failed, err:%v", err)
		return
	}
	theID, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		zlog.Errorf("get lastinsert ID failed, err:%v", err)
		return
	}
	zlog.Debugf("insert success, the id is %d.", theID)
	return
}

//GetUserInfo 获取用户详细信息
func GetUserInfo(id int) (userInfo *model.UserInfo, err error) {
	userInfo = &model.UserInfo{}
	sqlStr := "select user_id, user_name, avatar,create_time,mobile,mail,career,location, major,edu_background,working_age,introduce,has_edit from user where user_id=?"
	err = db.Get(userInfo, sqlStr, id)
	return
}

//EditUser 修改用户

//EditUserInfo 完善用户信息
func EditUserInfo(userInfo *model.UserInfo) (err error) {
	sqlStr := `
	UPDATE user SET  avatar=?,mobile=?,mail=?,career=?,location=?, major=?,edu_background=?,working_age=?,introduce=?,has_edit=?,update_time=?
	WHERE user_id=?
	`
	_, err = db.Exec(sqlStr, userInfo.Avatar, userInfo.Mobile, userInfo.Mail, userInfo.Career, userInfo.Location, userInfo.Major, userInfo.Edubackground, userInfo.Workingage, userInfo.Introduce, true, time.Now(), userInfo.UserID)

	return
}

//IsAdmin 用户是否为管理员权限
func IsAdmin(userID int) (b bool, err error) {
	sqlStr := "SELECT (SELECT u.role from user u where u.user_id=?)='admin'"

	err = db.Get(&b, sqlStr, userID)
	return
}

//GetSoftware (software_id) (model.Software)查看软件
//GetSoftwareList (user_id) ([]model.Software)查看软件列表
//GetSoftwareListByNum (user_id,page_num,page_size) ()查看软件列表
//AddSoftware (model.Software)新增软件
//UpdateSoftware (software_id)修改软件
//DelSoftware (software_id)删除软件
