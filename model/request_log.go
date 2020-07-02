package model

import "time"

//RequestLog 请求记录
type RequestLog struct {
	RequestID  int       `json:"request_id" db:"request_id" redis:"request_id"`
	CreateTime time.Time `json:"create_time" db:"create_time" redis:"create_time"`
	UserAgent  string    `json:"user_agent" db:"user_agent" redis:"user_agent"`
	URL        string    `json:"url" db:"url" redis:"url"`
	TokenInfo  string    `json:"token_info" db:"token_info" redis:"token_info"`
	Cookie     string    `json:"cookie" db:"cookie" redis:"cookie"`
	Method     string    `json:"method" db:"method" redis:"method"`
}
