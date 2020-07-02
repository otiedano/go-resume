package service

import (
	"encoding/json"
	"strconv"
	"sz_resume_202005/dao/db"
	"sz_resume_202005/dao/redisdb"
	"sz_resume_202005/model"
	"sz_resume_202005/utils"
	"sz_resume_202005/utils/setting"
	"sz_resume_202005/utils/zlog"
)

//CheckUserAuth 查看是否能登录
func CheckUserAuth(u *model.UserAuth) (ifExist bool, err error) {
	ifExist, err = db.CheckUserAuth(u)
	return
}

//GetUser 获取用户基本信息
func GetUser(u *model.UserAuth) (user *model.User, err error) {
	user, err = db.GetUser(u)
	return
}

//GetUserByID 通过ID查找用户

//EditUser 编辑用户信息-暂时不需要了

//AddUser 添加用户

//DelUser 删除用户

// //ExistByName 查看用户是否存在
// func ExistByName(userName string) (ifExist bool, err error) {
// 	ifExist, err = db.CheckUser(userName)
// 	return
// }

//ExistByID 通过ID检查用户是否存在

//GetUserInfo 查看用户详情
func GetUserInfo(userID int) (*model.UserInfo, error) {
	return db.GetUserInfo(userID)
}

//EditUserInfo 编辑用户信息
func EditUserInfo(userInfo *model.UserInfo) error {
	return db.EditUserInfo(userInfo)
}

//修改密码

//GenToken 生成token，并将User信息存入redis，有效期一周
func GenToken(u *model.User) (string, error) {
	byte, err := json.Marshal(*u)
	if err != nil {
		return "", err
	}
	t := utils.GenToken(u.UserName)
	err = redisdb.Set(strconv.Itoa(u.UserID), t, setting.AppSetting.TokenExpire)
	if err != nil {
		return "", err
	}
	err = redisdb.Set(t, string(byte), setting.AppSetting.TokenExpire)
	if err != nil {
		return "", err
	}
	return t, nil
}

//CheckToken 检查token
func CheckToken(s string) (string, bool, error) {
	return redisdb.Get(s)
}

//RemoveToken 移除对应用户id的token
func RemoveToken(id string) (err error) {
	token, exist, err := redisdb.Get(id)
	if err != nil {
		return
	}
	if !exist {
		return nil
	}

	//写一起怕的就是出现错误，id找不到token,先这么写
	_, err = redisdb.Del(token, id)
	if err != nil {
		return
	}
	zlog.Debugf("removeTOken,id:%v,token:%v", id, token)
	return nil
}

//IsAdmin 判断用户权限是否为管理员
func IsAdmin(userID int) (bool, error) {
	return db.IsAdmin(userID)
}
