package g

import (
	"sz_resume_202005/utils/e"

	"github.com/gin-gonic/gin"
)

//Gin 便于操作
type Gin struct {
	C *gin.Context
}

//Response 格式化的相应
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

//G 生成Gin对象
func G(c *gin.Context) *Gin {
	return &Gin{C: c}
}
