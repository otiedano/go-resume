package api

import (
	"net/http"
	"sz_resume_202005/model"
	"sz_resume_202005/service"
	"sz_resume_202005/utils/e"
	"sz_resume_202005/utils/g"
	"sz_resume_202005/utils/zlog"

	"github.com/gin-gonic/gin"
)

//AddSkill 添加工作经历
func AddSkill(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	zlog.Debugf("user:%+v,type:%T,UserID:%v", user, user, user.UserID)

	var skill model.Skill

	httpCode, errCode := g.BindAndValid(c, &skill)
	if errCode != e.SUCCESS {
		g.G(c).Response(httpCode, errCode, nil)
		return
	}

	zlog.Debugf("skill:%#v\n", skill)

	err := service.AddSkill(user.UserID, &skill)
	if err != nil {
		zlog.Errorf("run service.AddSkill failed,err~:%v", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	zlog.Debug("AddSkill:", skill)

	skills, err := service.GetSkills(user.UserID)
	if err != nil {
		zlog.Errorf("get skills failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, gin.H{
		"status": "success",
		"skills": skills,
	})

}

//GetSkills 获取工作经历
func GetSkills(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	skills, err := service.GetSkills(user.UserID)
	if err != nil {
		zlog.Errorf("get skills failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, gin.H{
		"status": "success",
		"skills": skills,
	})

}

//EditSkill 编辑工作经历
func EditSkill(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	var skill model.Skill

	httpCode, errCode := g.BindAndValid(c, &skill)
	if errCode != e.SUCCESS {
		g.G(c).Response(httpCode, errCode, nil)
		return
	}

	zlog.Debugf("skill:%v", skill)
	b, err := service.ExistSkill(user.UserID, &skill)
	if err != nil {
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	if !b {
		zlog.Errorf("no exist record")
		g.G(c).Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	err = service.EditSkill(user.UserID, &skill)
	if err != nil {
		zlog.Errorf("service.EditSkill failed,err:%v", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	skills, err := service.GetSkills(user.UserID)
	if err != nil {
		zlog.Errorf("get skills failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, gin.H{
		"status": "success",
		"skills": skills,
	})

}

//DelSkills 删除工作经历
func DelSkills(c *gin.Context) {
	u := c.MustGet("user")
	zlog.Debugf("u:%+v,type:%T,%v", u, u)
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	var a struct {
		Skills []int `json:"skills"`
	}

	err := c.Bind(&a)
	if err != nil {
		return
	}
	err = service.DelSkills(user.UserID, a.Skills)
	if err != nil {
		zlog.Errorf("service.DelSkills failed,err:", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	skills, err := service.GetSkills(user.UserID)
	if err != nil {
		zlog.Errorf("get skills failed,err:%v\n", err)
		g.G(c).Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	g.G(c).Response(http.StatusOK, e.SUCCESS, gin.H{
		"status": "success",
		"skills": skills,
	})
}
