package middleware

import (
	"bytes"
	"io/ioutil"
	"sz_resume_202005/model"
	"sz_resume_202005/service"
	"sz_resume_202005/utils/zlog"

	"github.com/gin-gonic/gin"
)

//RequestInfo 记录请求信息
func RequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		zlog.Debug("RequestInfo")

		body, _ := ioutil.ReadAll(c.Request.Body)
		zlog.Debugf("---body/--- \r\n " + string(body))
		zlog.Debug(c.Request.Header.Get("Content-Type"))
		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		rl := model.RequestLog{}
		rl.UserAgent = c.Request.Header.Get("User-Agent")
		rl.Cookie = c.Request.Header.Get("Cookie")
		rl.URL = c.Request.URL.String()
		rl.Method = c.Request.Method
		s, _, err := service.CheckToken(c.Request.Header.Get("Authorization"))
		if err != nil {
			zlog.Error(err)
		} else {
			rl.TokenInfo = s
		}
		//url,tokeninfo,Cookie  User-Agent

		err = service.AddRequestLog(&rl)
		if err != nil {
			zlog.Error("service.AddRequestLog failed,err:%v", err)
		}

		c.Next()
	}
}
