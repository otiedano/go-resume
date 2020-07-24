package v1

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

// var articleCategories struct {
// 	model.ArticleCategories `json:"article_categories" `
// }

//GetArticleCategory 获取全部分类
func GetArticleCategory(c *gin.Context) {
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
	articleCategory, err := service.GetArticleCategories()
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, articleCategory)

}

//GetAllArticleCategory 获取全部分类
func GetAllArticleCategory(c *gin.Context) {
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
	articleCategory, err := service.GetAllArticleCategories()
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, articleCategory)

}

//AddArticleCategory 添加分类
func AddArticleCategory(c *gin.Context) {
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

	articleCategory := &model.ArticleCategory{}

	err = c.ShouldBind(articleCategory)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	valid := validation.Validation{}
	valid.Required(articleCategory.CategoryName, "name")
	valid.Required(articleCategory.CategoryNo, "no")
	//检查参数合法性
	ok, _ = valid.Valid(articleCategory)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	intNum, err := service.AddArticleCategory(articleCategory)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_ADD_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, gin.H{
		"article_category_id": intNum,
	})
}

//EditArticleCategory 编辑分类
func EditArticleCategory(c *gin.Context) {
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

	articleCategory := &model.ArticleCategory{}

	c.ShouldBind(articleCategory)

	valid := validation.Validation{}
	valid.Required(articleCategory.CategoryName, "name")
	valid.Required(articleCategory.CategoryNo, "no")
	valid.Required(articleCategory.CategoryID, "id")
	//检查参数合法性
	ok, _ = valid.Valid(articleCategory)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = service.EditArticleCategory(articleCategory)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_ADD_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}

//BatchEditArticleCategory 批量编辑分类
func BatchEditArticleCategory(c *gin.Context) {
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

	var acs model.ArticleCategories
	err = c.ShouldBind(&acs)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD, nil)
		return
	}
	valid := validation.Validation{}
	for _, v := range acs {
		valid.Required((*v).CategoryName, "name ")
		valid.Required((*v).CategoryNo, "no ")
		valid.Required((*v).CategoryID, "id ")
		//检查参数合法性
		ok, _ = valid.Valid(*v)

		if !ok {
			zlog.MutiErrors(valid.Errors)
			g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
			return
		}
	}

	err = service.EditArticleCategories(&acs)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)

}

//DelArticleCategory 删除分类
func DelArticleCategory(c *gin.Context) {
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

	var acs model.ArticleCategories
	err = c.ShouldBind(&acs)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD, nil)
		return
	}

	err = service.DelArticleCategories(&acs)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)

}
