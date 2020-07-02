package middleware

import (
	"sz_resume_202005/model"
	"sz_resume_202005/service"
	"sz_resume_202005/utils/zlog"

	"github.com/gin-gonic/gin"
)

//RequestInfo 记录请求信息
func RequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		rl := model.RequestLog{}
		rl.UserAgent = c.Request.Header.Get("User-Agent")
		rl.Cookie = c.Request.Header.Get("Cookie")
		rl.URL = c.Request.URL.String()
		rl.Method = c.Request.Method
		s, _, _ := service.CheckToken(c.Request.Header.Get("Authorization"))
		//url,tokeninfo,Cookie  User-Agent
		rl.TokenInfo = s
		err := service.AddRequestLog(&rl)
		if err != nil {
			zlog.Error("service.AddRequestLog failed,err:%v", err)
		}

		c.Next()
	}
}
