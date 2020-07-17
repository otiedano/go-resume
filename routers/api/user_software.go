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

//AddSoftware 添加工作经历
func AddSoftware(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	zlog.Debugf("user:%+v,type:%T,UserID:%v", user, user, user.UserID)

	var software model.Software

	httpCode, errCode := g.BindAndValid(c, &software)
	if errCode != e.SUCCESS {
		g.G(c).Response(httpCode, errCode, nil)
		return
	}

	zlog.Debugf("software:%#v\n", software)

	err := service.AddSoftware(user.UserID, &software)
	if err != nil {
		zlog.Errorf("run software.AddSoftware failed,err~:%v", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	zlog.Debug("AddSoftware:", software)

	softwares, err := service.GetSoftwares(user.UserID)
	if err != nil {
		zlog.Errorf("get softwares failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, softwares)

}

//GetSoftwares 获取工作经历
func GetSoftwares(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	softwares, err := service.GetSoftwares(user.UserID)
	if err != nil {
		zlog.Errorf("get softwares failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, softwares)

}

//GetSoftware 获取工作经历
func GetSoftware(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	id := c.Param("id")
	swID, err := strconv.Atoi(id)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	s := &model.Software{UserID: user.UserID, SoftwareID: swID}
	b, err := service.ExistSoftware(user.UserID, s)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	if !b {

		zlog.Errorf("skill record not exist")
		g.Response(http.StatusBadRequest, e.ERROR_RECORD_NOT_EXIST, nil)
		return

	}

	software, err := service.GetSoftware(user.UserID, s.SoftwareID)
	if err != nil {
		zlog.Errorf("get skill failed,err:%v\n", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, software)

}

//EditSoftware 编辑工作经历
func EditSoftware(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	var software model.Software

	httpCode, errCode := g.BindAndValid(c, &software)
	if errCode != e.SUCCESS {
		g.G(c).Response(httpCode, errCode, nil)
		return
	}

	zlog.Debugf("software:%v", software)
	b, err := service.ExistSoftware(user.UserID, &software)
	if err != nil {
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	if !b {
		zlog.Errorf("no exist record")
		g.G(c).Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	err = service.EditSoftware(user.UserID, &software)
	if err != nil {
		zlog.Errorf("service.EditSoftware failed,err:%v", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	softwares, err := service.GetSoftwares(user.UserID)
	if err != nil {
		zlog.Errorf("get softwares failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, softwares)

}

//DelSoftwares 删除工作经历
func DelSoftwares(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	var a struct {
		Softwares []int `json:"softwares"`
	}

	err := c.Bind(&a)
	if err != nil {
		return
	}
	err = service.DelSoftwares(user.UserID, a.Softwares)
	if err != nil {
		zlog.Errorf("service.DelSoftwares failed,err:", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	softwares, err := service.GetSoftwares(user.UserID)
	if err != nil {
		zlog.Errorf("get softwares failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, gin.H{
		"status":    "success",
		"softwares": softwares,
	})
}
