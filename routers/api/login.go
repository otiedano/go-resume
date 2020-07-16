package api

import (
	"net/http"
	"strconv"
	"sz_resume_202005/model"
	"sz_resume_202005/service"
	"sz_resume_202005/utils/e"
	"sz_resume_202005/utils/g"
	"sz_resume_202005/utils/zlog"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//Login 验证用户登录,返回token
func Login(c *gin.Context) {
	zlog.Debug("调用 loginhandler")
	g := g.G(c)
	valid := validation.Validation{}

	userAuth := &model.UserAuth{}
	err := c.ShouldBind(userAuth)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	//检查参数合法性
	ok, _ := valid.Valid(userAuth)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	isExist, err := service.CheckUserAuth(userAuth)
	if err != nil {

		g.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !isExist {
		g.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	user, err := service.GetUser(userAuth)
	zlog.Debugw("service.GetUser", "user", user, "err", err)
	if err != nil {
		g.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	var token string
	token, err = service.GenToken(user)
	if err != nil {
		g.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	//t := time.Now()

	//c.SetCookie("token", token, setting.AppSetting.TokenExpire, "/", setting.AppSetting.Domain, false, true)
	//c.SetCookie("timestamp", strconv.FormatInt(t.Unix(), 10), setting.AppSetting.TokenExpire, "/", setting.AppSetting.Domain, false, true)
	scook, _ := c.Cookie("token")
	zlog.Debug("scook", scook)

	g.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})

}

//Logout 退出登录
func Logout(c *gin.Context) {
	u := c.MustGet("user").(*model.User)
	err := service.RemoveToken(strconv.Itoa(u.UserID))
	if err != nil {
		zlog.Debugf("LoginOut failed,err:%v", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, nil)

	//c.SetCookie("token", "", -1, "/", setting.AppSetting.Domain, false, true)
	//c.SetCookie("timestamp", "", -1, "/", setting.AppSetting.Domain, false, true)

}

//GetSkillAndSoftware 获取用户技能和软件
func GetSkillAndSoftware(c *gin.Context) {

}

//修改密码
