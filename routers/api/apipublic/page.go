package apipublic

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

//GetUserInfoCpt 首页内容
func GetUserInfoCpt(c *gin.Context) {
	g := g.G(c)
	user := &model.UserInfoCpt{}
	u, err := service.GetUserInfo(1)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORD, nil)
		return
	}
	exp, err := service.GetExperience(1)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORD, nil)
		return
	}
	user.UserInfo = *u
	user.Experiences = exp
	g.Response(http.StatusOK, e.SUCCESS, user)

}

//GetTec 软件和技能页
func GetTec(c *gin.Context) {
	g := g.G(c)
	skills, err := service.GetSkills(1)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	sws, err := service.GetSoftwares(1)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, gin.H{
		"softwares": sws,
		"skills":    skills,
	})
}

//GetArticles 获取文章列表
func GetArticles(c *gin.Context) {

	g := g.G(c)
	//获取分类和page
	p := c.DefaultQuery("page", "1")

	cg := c.DefaultQuery("category", "0")

	page, err := strconv.Atoi(p)
	if err != nil {
		zlog.Errorw("参数转换失败", "err", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	category, err := strconv.Atoi(cg)
	if err != nil {
		zlog.Errorw("参数转换失败", "err", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	articles, err := service.GetArticlesByCategory(category, page)

	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, articles)

}

//GetArticleByID 通过ID获取文章详情
func GetArticleByID(c *gin.Context) {

	g := g.G(c)
	//获取分类和page
	pid := c.Param("id")

	id, err := strconv.Atoi(pid)
	if err != nil {
		zlog.Errorw("参数转换失败", "err", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	service.CountArticle(id)
	b, err := service.ExistArticleByID(id)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	if !b {
		zlog.Errorw("文章记录不存在", "id", id)
		g.Response(http.StatusBadRequest, e.ERROR_RECORD_NOT_EXIST, nil)
		return
	}
	article, err := service.GetArticle(id)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, article)
}

//GetWork 获取作品列表
func GetWork(c *gin.Context) {
	g := g.G(c)
	//获取分类和page
	p := c.DefaultQuery("page", "1")

	t := c.DefaultQuery("tag", "0")

	page, err := strconv.Atoi(p)
	if err != nil {
		zlog.Errorw("参数转换失败", "err", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	tag, err := strconv.Atoi(t)
	if err != nil {
		zlog.Errorw("参数转换失败", "err", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	works, err := service.GetWorksByTag(tag, page)

	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, works)
}

//GetWorkByID 通过ID获取作品详情
func GetWorkByID(c *gin.Context) {
	g := g.G(c)
	//获取分类和page
	pid := c.Param("id")
	zlog.Debug("id:", pid)
	id, err := strconv.Atoi(pid)
	if err != nil {
		zlog.Errorw("参数转换失败", "err", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	service.CountWork(id)

	b, err := service.ExistWorkByID(id)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	if !b {
		zlog.Errorw("文章记录不存在", "id", id)
		g.Response(http.StatusBadRequest, e.ERROR_RECORD_NOT_EXIST, nil)
		return
	}
	article, err := service.GetWork(id)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, article)
}
