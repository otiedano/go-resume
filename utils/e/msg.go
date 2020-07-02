package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	INTERNALERROR:  "fail",
	INVALID_PARAMS: "请求参数错误",
	UNAUTHORIZED:   "没有访问权限,请登录或向管理员申请权限",

	ERROR_ADD_RECORD:       "新增记录失败",
	ERROR_RECORD_NOT_EXIST: "该记录不存在",
	ERROR_EDIT_RECORD:      "修改记录失败",
	ERROR_EDIT_RECORDS:     "修改多条记录失败",
	ERROR_GET_RECORD:       "获取记录失败",
	ERROR_GET_RECORDS:      "获取多条记录失败",
	ERROR_DEL_RECORDS:      "删除记录失败",
	ERROR_RECORD_EXIST:     "记录已存在",
	ERROR_COUNT_RECORD:     "统计记录失败",
	ERROR_UPDATE_RECORD:    "更新访问信息失败",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "用户名或密码错误",

	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[INTERNALERROR]
}
