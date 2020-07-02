package g

import (
	"net/http"
	"sz_resume_202005/utils/e"
	"sz_resume_202005/utils/zlog"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.INTERNALERROR
	}
	if !check {

		zlog.MutiErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
