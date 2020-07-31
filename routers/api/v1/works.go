package v1

import (
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

type checkWork struct {
	Status int   `json:"status"`
	Works  []int `json:"works"`
}

//GetWork 根据用户获取作品列表
func GetWork(c *gin.Context) {
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

	works, err := service.GetWorksByAuthor(user.UserID, page)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_GET_RECORDS, nil)
		return
	}
	total, err := service.TotalWorkByAuthor(user.UserID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_COUNT_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, gin.H{
		"total":   total,
		"current": page,
		"size":    setting.PageSize,
		"works":   works,
	})
}

//GetWorkDetail 获取作品详情
func GetWorkDetail(c *gin.Context) {

	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	id := c.Param("id")
	workID, err := strconv.Atoi(id)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	b, err := service.ExistWorkByAuth(user.UserID, workID)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	if !b {

		zlog.Errorf("workID record not exist")
		g.Response(http.StatusBadRequest, e.ERROR_RECORD_NOT_EXIST, nil)
		return

	}
	ad, err := service.GetPWork(user.UserID, workID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, ad)

}

//AddWork 添加作品
func AddWork(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}

	work := &model.WorkDetail{}
	err := c.ShouldBind(work)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	valid := validation.Validation{}
	valid.Required(work.Title, "title ")
	valid.Required(work.CoverImg, "cover_img ")
	valid.Required(work.StartTime, "start_time ")
	valid.Required(work.EndTime, "end_time ")
	valid.Required(work.Content, "content ")
	valid.Required(work.Img, "img  ")
	valid.Required(work.IfHorizons, "if_horizons ")
	valid.Required(work.WorkNO, "work_no ")
	if len(work.WorkImgs) > 0 {
		for _, v := range work.WorkImgs {
			valid.Required(v.ImgNo, "img_no ")
			valid.Required(v.ImgPath, "img_path ")
		}
	}

	//检查参数合法性
	ok, _ = valid.Valid(work)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	if len(work.Tags) > 0 {
		for _, v := range work.Tags {
			valid.Required(v.TagID, "tag_id ")
			b, err := service.ExistWorkTag(v.TagID)
			if err != nil {
				zlog.Error(err)
				g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
				return
			}
			if !b {
				zlog.Error("v.TagID not exist")
				g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
				return
			}
		}

	}
	intNum, err := service.AddWork(user.UserID, work)
	if err != nil {
		zlog.Errorf("service.AddArticle failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.ERROR_ADD_RECORD, nil)

	}
	g.Response(http.StatusOK, e.SUCCESS, gin.H{"work_id": intNum})
}

//EditWork 编辑作品
func EditWork(c *gin.Context) {
	g := g.G(c)
	u := c.MustGet("user")
	user, ok := u.(*model.User)
	if !ok {
		zlog.Errorf("user assertion error.\n")
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	id := c.Param("id")
	workID, err := strconv.Atoi(id)
	work := &model.WorkDetail{}
	err = c.ShouldBind(work)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	work.WorkID = workID
	valid := validation.Validation{}
	valid.Required(work.Title, "title ")
	valid.Required(work.WorkID, "work_id ")
	valid.Required(work.CoverImg, "cover_imt ")
	valid.Required(work.StartTime, "start_time ")
	valid.Required(work.EndTime, "end_time ")
	valid.Required(work.Content, "content ")
	valid.Required(work.Img, "img  ")
	valid.Required(work.IfHorizons, "if_horizons ")
	valid.Required(work.WorkNO, "work_no ")
	if len(work.WorkImgs) > 0 {
		for _, v := range work.WorkImgs {
			valid.Required(v.ImgNo, "img_no ")
			valid.Required(v.ImgPath, "img_path ")
		}
	}

	//检查参数合法性
	ok, _ = valid.Valid(work)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	if len(work.Tags) > 0 {
		for _, v := range work.Tags {
			valid.Required(v.TagID, "tag_id ")
			b, err := service.ExistWorkTag(v.TagID)
			if err != nil {
				zlog.Error(err)
				g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
				return
			}
			if !b {
				zlog.Error("v.TagID not exist")
				g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
				return
			}
		}

	}
	b, err := service.ExistWorkByAuth(user.UserID, workID)
	if err != nil {
		zlog.Errorf("service.ExistWorkByAuth failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	if !b {
		zlog.Error("workID is not exist")
		g.Response(http.StatusBadRequest, e.ERROR_RECORD_NOT_EXIST, nil)
		return
	}
	err = service.EditWork(user.UserID, work)
	if err != nil {
		zlog.Errorf("service.AddArticle failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.ERROR_ADD_RECORD, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}

//DelWork 删除作品
func DelWork(c *gin.Context) {
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

	zlog.Debugf("ids:%v,len:%v\n", ids, len(ids))
	err = service.DelWorks(user.UserID, ids)
	if err != nil {
		zlog.Errorf("service.DelWorks failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}

//RACheckWork 管理员审核作品
func RACheckWork(c *gin.Context) {
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

	w := &checkWork{}
	err = c.ShouldBind(w)
	if err != nil {
		zlog.Errorf("ShouldBind failed,err:%v", err)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	valid := validation.Validation{}

	valid.Range(w.Status, 0, 100, "status range ")
	valid.Required(w.Works, "work ")

	//检查参数合法性
	ok, _ = valid.Valid(w)

	if !ok {
		zlog.MutiErrors(valid.Errors)
		g.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	err = service.RACheckWorks(w.Works, w.Status)
	if err != nil {
		zlog.Errorf("service.RACheckArticles failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}

//RAGetWork 管理员根据状态读取作品列表
func RAGetWork(c *gin.Context) {
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

	total, err := service.RATotalWork(status)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_INIT_PARAM, nil)
		return
	}
	works, err := service.RAGetAllWorksByStatus(page, status)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.ERROR_INIT_PARAM, nil)
		return
	}

	g.Response(http.StatusOK, e.SUCCESS, gin.H{
		"total":   total,
		"current": page,
		"size":    setting.PageSize,
		"works":   works,
	})
}

//RAGetWorkDetail 管理员读取作品详情
func RAGetWorkDetail(c *gin.Context) {
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
	workID, err := strconv.Atoi(id)
	if err != nil {

		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return

	}
	b, err = service.RAExistWork(workID)
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
	workDetail, err := service.RAGetWork(workID)
	if err != nil {
		zlog.Error(err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, workDetail)
}

//RADelWork 管理员删除作品
func RADelWork(c *gin.Context) {
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

	err = service.RADelWorksFE(ids)
	if err != nil {
		zlog.Errorf("service.RADelWorksFE failed,err:%v", err)
		g.Response(http.StatusInternalServerError, e.INTERNALERROR, nil)
		return
	}
	g.Response(http.StatusOK, e.SUCCESS, nil)
}
