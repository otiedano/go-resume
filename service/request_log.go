package service

import (
	"sz_resume_202005/dao/db"
	"sz_resume_202005/model"
	"sz_resume_202005/utils/setting"
)

//AddRequestLog 记录请求信息
func AddRequestLog(rl *model.RequestLog) (err error) {
	return db.AddRequestLog(rl)
}

//GetRequestLog 按分页读取请求记录
func GetRequestLog(page int) {
	offset := (page - 1) * setting.PageSize
	db.GetRequestLog(offset, setting.PageSize)
}
