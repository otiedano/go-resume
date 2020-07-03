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

//GetWorkTag 获取全部标签
func GetWorkTag(c *gin.Context) {
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
	workTag, err := service.GetWorkTags()
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, workTag)

}

//AddWorkTag 添加标签
func AddWorkTag(c *gin.Context) {

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

	workTag := &model.WorkTag{}

	err = c.ShouldBind(workTag)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	valid := validation.Validation{}
	valid.Required(workTag.TagName, "name")
	valid.Required(workTag.TagNO, "no")
	//检查参数合法性
	ok, _ = valid.Valid(workTag)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	intNum, err := service.AddWorkTag(workTag)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_ADD_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, gin.H{
		"work_tag_id": intNum,
	})
}

//EditWorkTag 编辑标签
func EditWorkTag(c *gin.Context) {

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

	workTag := &model.WorkTag{}

	c.ShouldBind(workTag)

	valid := validation.Validation{}
	valid.Required(workTag.TagName, "name")
	valid.Required(workTag.TagNO, "no")
	valid.Required(workTag.TagID, "id")
	//检查参数合法性
	ok, _ = valid.Valid(workTag)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	err = service.EditWorkTag(workTag)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_ADD_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}

//BatchEditWorkTag 批量编辑标签
func BatchEditWorkTag(c *gin.Context) {

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

	var wts model.WorkTags
	err = c.ShouldBind(&wts)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD, nil)
		return
	}
	valid := validation.Validation{}
	for _, v := range wts {
		valid.Required((*v).TagName, "name ")
		valid.Required((*v).TagNO, "no ")
		valid.Required((*v).TagID, "id ")
		//检查参数合法性
		ok, _ = valid.Valid(*v)

		if !ok {
			zlog.MutiErrors(valid.Errors)
			g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
			return
		}
	}

	err = service.EditWorkTags(&wts)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)

}

//DelWorkTag 删除标签
func DelWorkTag(c *gin.Context) {

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

	var wts model.WorkTags
	err = c.ShouldBind(&wts)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD, nil)
		return
	}

	err = service.DelWorkTags(&wts)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_EDIT_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)

}
