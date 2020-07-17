package api

import (
	"net/http"
	"strconv"
	"sz_resume_202005/model"
	"sz_resume_202005/service"
	"sz_resume_202005/utils/e"
	"sz_resume_202005/utils/g"
	"sz_resume_202005/utils/zlog"

	"github.com/gin-gonic/gin"
)

//AddExperiences 添加工作经历
func AddExperiences(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	zlog.Debugf("user:%+v,type:%T,UserID:%v", user, user, user.UserID)

	var ep struct {
		model.Experiences `json:"experiences" `
	}

	err := c.BindJSON(&ep)
	if err != nil {
		zlog.Errorf("bind failed,err~:%v", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	zlog.Debugf("ep:%#v\n", ep)
	if len(ep.Experiences) > 0 {
		err = service.AddExperiences(user.UserID, ep.Experiences)
		if err != nil {
			zlog.Errorf("run service.AddExperiences failed,err~:%v", err)
			g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
			return
		}

	} else {
		zlog.Errorf("invalid params.\n")
		g.G(c).Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	zlog.Debug("addexperience:", ep)

	exps, err := service.GetExperiences(user.UserID)
	if err != nil {
		zlog.Errorf("get experience failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, gin.H{
		"status":     "success",
		"expriences": exps,
	})

}

//GetExperiences 获取工作经历
func GetExperiences(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	exps, err := service.GetExperiences(user.UserID)
	if err != nil {
		zlog.Errorf("get experience failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, exps)

}

//GetExperience 获取工作经历
func GetExperience(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g := g.G(c)
	id := c.Param("id")
	expID, err := strconv.Atoi(id)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	ex := &model.Experience{UserID: user.UserID, ExpID: expID}
	b, err := service.ExistExperience(user.UserID, ex)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	if !b {

		zlog.Errorf("experience record not exist")
		g.Response(http.StatusBadRequest, e.ERROR_RECORD_NOT_EXIST, nil)
		return

	}

	exp, err := service.GetExperience(user.UserID, ex.ExpID)
	if err != nil {
		zlog.Errorf("get experience failed,err:%v\n", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, exp)

}

//EditExperience 编辑工作经历
func EditExperience(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	var exp model.Experience
	err := c.ShouldBind(&exp)
	if err != nil {
		zlog.Errorf("bind exp failed,err:%v", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	zlog.Debugf("exp:%v", exp)
	b, err := service.ExistExperience(user.UserID, &exp)
	if err != nil {
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	if !b {
		zlog.Errorf("no exist record")
		g.G(c).Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	err = service.EditExperience(user.UserID, &exp)
	if err != nil {
		zlog.Errorf("service.EditExperience failed,err:%v", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	exps, err := service.GetExperiences(user.UserID)
	if err != nil {
		zlog.Errorf("get experience failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, gin.H{
		"status":      "success",
		"experiences": exps,
	})

}

//DelExperiences 删除工作经历
func DelExperiences(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	var ids []int
	err := c.ShouldBind(&ids)
	if err != nil {
		zlog.Errorf("ShouldBind failed,err:%v", err)
		g.G(c).Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	zlog.Debugf("ids:%v", ids)
	zlog.Debugf("%T", ids)

	if err != nil {
		return
	}
	if l := len(ids); l <= 0 {
		zlog.Errorf("len(a)<=0")
		g.G(c).Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = service.DelExperiences(user.UserID, ids)
	if err != nil {
		zlog.Errorf("service.DelExperiences failed,err:", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	exps, err := service.GetExperiences(user.UserID)
	if err != nil {
		zlog.Errorf("get experience failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, gin.H{
		"status":      "success",
		"experiences": exps,
	})
}
