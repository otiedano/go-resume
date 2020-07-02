package db

import "sz_resume_202005/model"

//AddRequestLog 记录请求信息
func AddRequestLog(rl *model.RequestLog) (err error) {
	sqlStr := "INSERT INTO request_log (url,user_agent,token_info,cookie,method) VALUES (?,?,?,?,?)"
	_, err = db.Exec(sqlStr, rl.URL, rl.UserAgent, rl.TokenInfo, rl.Cookie, rl.Method)
	return
}

//GetRequestLog 读取请求信息
func GetRequestLog(offset, limit int) (rls []*model.RequestLog, err error) {
	sqlStr := "SELECT request_id,create_time,user_agent,url,token_info,cookie,method from request_log LIMIT ? OFFSET ?"
	err = db.Select(&rls, sqlStr, limit, offset)
	return
}
