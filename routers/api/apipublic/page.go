package apipublic

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"sz_resume_202005/model"
	"sz_resume_202005/service"
	"sz_resume_202005/utils/e"
	"sz_resume_202005/utils/g"
	"sz_resume_202005/utils/setting"
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

	if strings.Index(u.Avatar, "http://") != 0 && strings.Index(u.Avatar, "https://") != 0 {
		if strings.Index(u.Avatar, "/") == 0 {
			u.Avatar = setting.AppSetting.PrefixURL + u.Avatar
		} else {
			u.Avatar = setting.AppSetting.PrefixURL + "/" + u.Avatar
		}

	}

	exp, err := service.GetExperiences(1)
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
	for _, v := range skills {
		if strings.Index(v.Img, "http://") != 0 && strings.Index(v.Img, "https://") != 0 {
			if strings.Index(v.Img, "/") == 0 {
				v.Img = setting.AppSetting.PrefixURL + v.Img
			} else {
				v.Img = setting.AppSetting.PrefixURL + "/" + v.Img
			}

		}

	}
	sws, err := service.GetSoftwares(1)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	for _, v := range sws {

		if strings.Index(v.Img, "http://") != 0 && strings.Index(v.Img, "https://") != 0 {
			if strings.Index(v.Img, "/") == 0 {
				v.Img = setting.AppSetting.PrefixURL + v.Img
			} else {
				v.Img = setting.AppSetting.PrefixURL + "/" + v.Img
			}

		}

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

	total, err := service.TotalArticleByCategory(category)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_COUNT_RECORD, nil)
		return
	}
	if len(articles) == 0 && total > 0 && page > 1 {

		page = int(math.Ceil(float64(total) / float64(setting.PageSize)))
		articles, err = service.GetArticlesByCategory(category, page)

		if err != nil {
			zlog.Error(err)
			g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
			return
		}
	}

	for _, v := range articles {

		if strings.Index(v.Img, "http://") != 0 && strings.Index(v.Img, "https://") != 0 {
			if strings.Index(v.Img, "/") == 0 {
				v.Img = setting.AppSetting.PrefixURL + v.Img
			} else {
				v.Img = setting.AppSetting.PrefixURL + "/" + v.Img
			}

		}

		if strings.Index(v.Avatar, "http://") != 0 && strings.Index(v.Avatar, "https://") != 0 {
			if strings.Index(v.Avatar, "/") == 0 {
				v.Avatar = setting.AppSetting.PrefixURL + v.Avatar
			} else {
				v.Avatar = setting.AppSetting.PrefixURL + "/" + v.Avatar
			}

		}

	}

	g.Response(http.StatusOK, e.SUCCESS, gin.H{
		"total":    total,
		"current":  page,
		"size":     setting.PageSize,
		"articles": articles,
	})

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

	// article.Avatar = setting.AppSetting.PrefixURL + article.Avatar
	// article.Img = setting.AppSetting.PrefixURL + article.Img

	if strings.Index(article.Avatar, "http://") != 0 && strings.Index(article.Avatar, "https://") != 0 {
		if strings.Index(article.Avatar, "/") == 0 {
			article.Avatar = setting.AppSetting.PrefixURL + article.Avatar
		} else {
			article.Avatar = setting.AppSetting.PrefixURL + "/" + article.Avatar
		}

	}
	if strings.Index(article.Img, "http://") != 0 && strings.Index(article.Img, "https://") != 0 {
		if strings.Index(article.Img, "/") == 0 {
			article.Img = setting.AppSetting.PrefixURL + article.Img
		} else {
			article.Img = setting.AppSetting.PrefixURL + "/" + article.Img
		}

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

	total, err := service.TotalWorkByTag(tag)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_COUNT_RECORD, nil)
		return
	}

	if len(works) == 0 && total > 0 && page > 1 {

		page = int(math.Ceil(float64(total) / float64(setting.PageSize)))
		works, err = service.GetWorksByTag(tag, page)

		if err != nil {
			zlog.Error(err)
			g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
			return
		}
	}
	// for _, v := range works {
	// 	v.CoverImg = setting.AppSetting.PrefixURL + v.CoverImg
	// 	v.Avatar = setting.AppSetting.PrefixURL + v.Avatar

	// }

	for _, v := range works {

		if strings.Index(v.CoverImg, "http://") != 0 && strings.Index(v.CoverImg, "https://") != 0 {
			if strings.Index(v.CoverImg, "/") == 0 {
				v.CoverImg = setting.AppSetting.PrefixURL + v.CoverImg
			} else {
				v.CoverImg = setting.AppSetting.PrefixURL + "/" + v.CoverImg
			}

		}

		if strings.Index(v.Avatar, "http://") != 0 && strings.Index(v.Avatar, "https://") != 0 {
			if strings.Index(v.Avatar, "/") == 0 {
				v.Avatar = setting.AppSetting.PrefixURL + v.Avatar
			} else {
				v.Avatar = setting.AppSetting.PrefixURL + "/" + v.Avatar
			}

		}

	}

	g.Response(http.StatusOK, e.SUCCESS, gin.H{
		"total":   total,
		"current": page,
		"size":    setting.PageSize,
		"works":   works,
	})

}

//GetAllWork 获取所有作品列表
func GetAllWork(c *gin.Context) {
	g := g.G(c)
	//获取分类和page

	works, err := service.GetWorksNoLimit()

	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}

	total, err := service.TotalWorkByTag(0)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_COUNT_RECORD, nil)
		return
	}

	// for _, v := range works {
	// 	v.CoverImg = setting.AppSetting.PrefixURL + v.CoverImg
	// 	v.Avatar = setting.AppSetting.PrefixURL + v.Avatar

	// }

	for _, v := range works {

		if strings.Index(v.CoverImg, "http://") != 0 && strings.Index(v.CoverImg, "https://") != 0 {
			if strings.Index(v.CoverImg, "/") == 0 {
				v.CoverImg = setting.AppSetting.PrefixURL + v.CoverImg
			} else {
				v.CoverImg = setting.AppSetting.PrefixURL + "/" + v.CoverImg
			}

		}

		if strings.Index(v.Avatar, "http://") != 0 && strings.Index(v.Avatar, "https://") != 0 {
			if strings.Index(v.Avatar, "/") == 0 {
				v.Avatar = setting.AppSetting.PrefixURL + v.Avatar
			} else {
				v.Avatar = setting.AppSetting.PrefixURL + "/" + v.Avatar
			}

		}

	}

	g.Response(http.StatusOK, e.SUCCESS, gin.H{
		"total": total,

		"works": works,
	})

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
	work, err := service.GetWork(id)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}

	// work.Avatar = setting.AppSetting.PrefixURL + work.Avatar
	// work.Img = setting.AppSetting.PrefixURL + work.Img
	// work.CoverImg = setting.AppSetting.PrefixURL + work.CoverImg

	if strings.Index(work.CoverImg, "http://") != 0 && strings.Index(work.CoverImg, "https://") != 0 {
		if strings.Index(work.CoverImg, "/") == 0 {
			work.CoverImg = setting.AppSetting.PrefixURL + work.CoverImg
		} else {
			work.CoverImg = setting.AppSetting.PrefixURL + "/" + work.CoverImg
		}

	}
	if strings.Index(work.Img, "http://") != 0 && strings.Index(work.Img, "https://") != 0 {
		if strings.Index(work.Img, "/") == 0 {
			work.Img = setting.AppSetting.PrefixURL + work.Img
		} else {
			work.Img = setting.AppSetting.PrefixURL + "/" + work.Img
		}

	}
	if strings.Index(work.Avatar, "http://") != 0 && strings.Index(work.Avatar, "https://") != 0 {
		if strings.Index(work.Avatar, "/") == 0 {
			work.Avatar = setting.AppSetting.PrefixURL + work.Avatar
		} else {
			work.Avatar = setting.AppSetting.PrefixURL + "/" + work.Avatar
		}

	}

	for _, v := range work.WorkImgs {
		if strings.Index(v.ImgPath, "http://") != 0 && strings.Index(v.ImgPath, "https://") != 0 {
			if strings.Index(v.ImgPath, "/") == 0 {
				v.ImgPath = setting.AppSetting.PrefixURL + v.ImgPath
			} else {
				v.ImgPath = setting.AppSetting.PrefixURL + "/" + v.ImgPath
			}

		}
	}

	g.Response(http.StatusOK, e.SUCCESS, work)
}
