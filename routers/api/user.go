package api

import (
	"net/http"
	"sz_resume_202005/model"
	"sz_resume_202005/service"
	"sz_resume_202005/utils/e"
	"sz_resume_202005/utils/g"
	"sz_resume_202005/utils/zlog"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user").(*model.User)
	userInfo, err := service.GetUserInfo(u.UserID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, userInfo)
}

//EditUserInfo 编辑用户信息
func EditUserInfo(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user").(*model.User)
	var userInfo model.UserInfo

	c.BindJSON(&userInfo)
	userInfo.UserID = u.UserID
	valid := validation.Validation{}
	valid.Email(userInfo.Mail, "email")

	ok, _ := valid.Valid(userInfo)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	zlog.Debugf("userInfo:%+v", userInfo)
	err := service.EditUserInfo(&userInfo)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}
