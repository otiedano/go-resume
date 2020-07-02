package service

import (
	"fmt"
	"sz_resume_202005/dao/db"
	"sz_resume_202005/dao/redisdb"
	"sz_resume_202005/model"
	"sz_resume_202005/utils/setting"
	"sz_resume_202005/utils/zlog"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(t *testing.T) {

	setting.Init()
	zlog.Init()
	defer zlog.Sync()
	redisdb.Init()
	db.Init()
	zlog.Debug("test~")
	gin.SetMode(gin.TestMode)
}
func TestGetUser(t *testing.T) {
	u := model.UserAuth{
		UserName: "tiedan",
		Password: "imtiedan",
	}
	user, err := GetUser(&u)
	if err != nil {
		t.Error("err:", err, "\n")
	}
	fmt.Printf("user:%#v\n", user)
}
func TestRemoveToken(t *testing.T) {
	err := RemoveToken("1")
	if err != nil {
		t.Error("err:", err)
	}
}
