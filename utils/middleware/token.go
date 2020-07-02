package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sz_resume_202005/model"
	"sz_resume_202005/service"
	"sz_resume_202005/utils/e"
	"sz_resume_202005/utils/g"
	"sz_resume_202005/utils/zlog"

	"github.com/gin-gonic/gin"
)

// CheckToken 检查请求是否携带token
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		zlog.Debug("checkToken")
		var (
			t     model.UserToken
			val   string
			err   error
			exist bool
		)

		t.Token = c.Request.Header.Get("Authorization")

		if len(t.Token) > 0 {
			val, exist, err = service.CheckToken(t.Token)
			fmt.Printf("err:%v\n", err)
			fmt.Printf("exist:%v\n", exist)
			fmt.Printf("val:%+v\n", val)
			if err != nil {

				g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
				c.Abort()
				return
			}
			if !exist {

				g.G(c).Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
				c.Abort()
				return

			}

		} else {
			err = c.BindJSON(&t)
			if err != nil {
				g.G(c).Response(http.StatusBadRequest, e.INTERNALERROR, nil)
				c.Abort()
				return
			}
			val, exist, err = service.CheckToken(t.Token)
			if err != nil {

				g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
				c.Abort()
				return
			}
			if !exist {

				g.G(c).Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
				c.Abort()
				return

			}
		}

		var user model.User
		err = json.Unmarshal([]byte(val), &user)
		zlog.Debugf("user:%+v", user)
		if err != nil {
			g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
			c.Abort()
			return
		}
		c.Set("user", &user)
		zlog.Debugw("checktoken通过", "user", user)
		c.Next()

	}
}
