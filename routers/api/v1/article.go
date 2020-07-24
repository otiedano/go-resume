package v1

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sz_resume_202005/model"
	"sz_resume_202005/service"
	"sz_resume_202005/utils/e"
	"sz_resume_202005/utils/g"
	"sz_resume_202005/utils/setting"
	"sz_resume_202005/utils/zlog"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type checkArticle struct {
	Status   int   `json:"status"`
	Articles []int `json:"articles"`
}

//GetArticle 根据用户获取文章列表
func GetArticle(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	pg := c.DefaultQuery("page", "1")

	page, err := strconv.Atoi(pg)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_INIT_PARAM, nil)
		return
	}

	articles, err := service.GetArticlesByAuthor(user.UserID, page)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	total, err := service.TotalArticleByAuthor(user.UserID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_COUNT_RECORD, nil)
		return
	}
	if len(articles) == 0 && total > 0 && page > 1 {

		page = int(math.Ceil(float64(total) / float64(setting.PageSize)))
		articles, err = service.GetArticlesByAuthor(user.UserID, page)
		if err != nil {
			zlog.Error(err)
			g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
			return
		}
	}
	g.Response(http.StatusOK, e.SUCCESS, gin.H{
		"total":    total,
		"current":  page,
		"size":     setting.PageSize,
		"articles": articles,
	})
}

//GetArticleDetail 获取文章详情
func GetArticleDetail(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	id := c.Param("id")
	articleID, err := strconv.Atoi(id)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	b, err := service.ExistArticleByAuth(user.UserID, articleID)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	if !b {

		zlog.Errorf("articleID record not exist")
		g.Response(http.StatusBadRequest, e.ERROR_RECORD_NOT_EXIST, nil)
		return

	}
	ad, err := service.GetPArticle(user.UserID, articleID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, ad)
}

//AddArticle 添加文章
func AddArticle(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	article := &model.ArticleDetail{}
	err := c.ShouldBind(article)
	if err != nil {
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}
	valid := validation.Validation{}
	valid.Required(article.Title, "title ")
	valid.Required(article.Content, "content ")
	valid.Required(article.Img, "img  ")
	valid.Required(article.CategoryID, "categoryID ")
	//检查参数合法性
	ok, _ = valid.Valid(article)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	b, err := service.ExistArticleCategory(article.CategoryID)
	if err != nil {
		zlog.Errorf("service.ExistArticleCategory failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	if !b {
		zlog.Errorf("CategoryID not exist")
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	intNum, err := service.AddArticle(user.UserID, article)
	if err != nil {
		zlog.Errorf("service.AddArticle failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.ERROR_ADD_RECORD, nil)

	}
	g.Response(http.StatusOK, e.SUCCESS, gin.H{"article_id": intNum})
}

//EditArticle 编辑文章
func EditArticle(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	id := c.Param("id")

	articleID, err := strconv.Atoi(id)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}

	article := &model.ArticleDetail{}
	err = c.ShouldBind(article)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	article.ArticleID = articleID
	fmt.Printf("article:%v\n", article)
	valid := validation.Validation{}
	valid.Required(article.ArticleID, "articleID ")
	valid.Required(article.Title, "title ")
	valid.Required(article.Content, "content ")
	valid.Required(article.Img, "img  ")
	valid.Required(article.CategoryID, "categoryID ")
	//检查参数合法性
	ok, _ = valid.Valid(article)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	b, err := service.ExistArticleByAuth(user.UserID, article.ArticleID)
	if err != nil {
		zlog.Errorf("service.ExistArticleByAuth failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	if !b {
		zlog.Errorf("ArticleID not exist")
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	b, err = service.ExistArticleCategory(article.CategoryID)
	if err != nil {
		zlog.Errorf("service.ExistArticleCategory failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	if !b {
		zlog.Errorf("CategoryID not exist")
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = service.EditArticle(user.UserID, article)
	if err != nil {
		zlog.Errorf("service.AddArticle failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.ERROR_ADD_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}

//DelArticle 删除文章
func DelArticle(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	var ids []int
	err := c.ShouldBind(&ids)
	if err != nil {
		zlog.Errorf("ShouldBind failed,err:%v", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	if len(ids) == 0 {
		zlog.Errorf("request param is non")
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	fmt.Printf("ids:%v,len:%v\n", ids, len(ids))
	err = service.DelArticles(user.UserID, ids)
	if err != nil {
		zlog.Errorf("service.DelArticles failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}

//RACheckArticle 管理员审核文章
func RACheckArticle(c *gin.Context) {

	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	b, err := service.IsAdmin(user.UserID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_ROLE, nil)
		return
	}

	if !b {
		zlog.Error(e.GetMsg(e.UNAUTHORIZED))
		g.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
		return
	}

	ca := &checkArticle{}
	err = c.ShouldBind(ca)
	if err != nil {
		zlog.Errorf("ShouldBind failed,err:%v", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	valid := validation.Validation{}

	valid.Range(ca.Status, 0, 100, "status range ")
	valid.Required(ca.Articles, "articles")

	//检查参数合法性
	ok, _ = valid.Valid(ca)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	err = service.RACheckArticles(ca.Articles, ca.Status)
	if err != nil {
		zlog.Errorf("service.RACheckArticles failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}

//RAGetArticle 管理员根据状态读取文章列表
func RAGetArticle(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	b, err := service.IsAdmin(user.UserID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_ROLE, nil)
		return
	}

	if !b {
		zlog.Error(e.GetMsg(e.UNAUTHORIZED))
		g.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
		return
	}

	pg := c.DefaultQuery("page", "1")

	page, err := strconv.Atoi(pg)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_INIT_PARAM, nil)
		return
	}
	st := c.DefaultQuery("status", "-1")

	status, err := strconv.Atoi(st)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_INIT_PARAM, nil)
		return
	}

	total, err := service.RATotalArticle(status)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_INIT_PARAM, nil)
		return
	}
	articles, err := service.RAGetAllArtilesByStatus(page, status)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_INIT_PARAM, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, gin.H{
		"total":    total,
		"current":  page,
		"size":     setting.PageSize,
		"articles": articles,
	})
}

//RAGetArticleDetail 管理员读取文章详情
func RAGetArticleDetail(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	b, err := service.IsAdmin(user.UserID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_ROLE, nil)
		return
	}

	if !b {
		zlog.Error(e.GetMsg(e.UNAUTHORIZED))
		g.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
		return
	}

	id := c.Param("id")
	articleID, err := strconv.Atoi(id)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	b, err = service.RAExistArticle(articleID)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	if !b {

		zlog.Errorf("articleID record not exist")
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return

	}
	ad, err := service.RAGetArticle(articleID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, ad)
}

//RADelArticle 管理员删除文章
func RADelArticle(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	b, err := service.IsAdmin(user.UserID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_ROLE, nil)
		return
	}

	if !b {
		zlog.Error(e.GetMsg(e.UNAUTHORIZED))
		g.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
		return
	}

	var ids []int
	err = c.ShouldBind(&ids)
	if err != nil {
		zlog.Errorf("ShouldBind failed,err:%v", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	if len(ids) == 0 {
		zlog.Errorf("request param is non")
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	err = service.RADelArticlesFE(ids)
	if err != nil {
		zlog.Errorf("service.RADelArticlesFE failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}
