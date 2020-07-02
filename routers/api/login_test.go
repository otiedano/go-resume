package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sz_resume_202005/dao/db"
	"sz_resume_202005/dao/redisdb"
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
func TestLogin(t *testing.T) {
	c := http.Client{}
	rd := strings.NewReader("{\"username\":\"tiedan\",\"password\":\"imtiedan\"}")
	reqest, err := http.NewRequest("POST", "http://localhost:8083/api/login", rd)
	if err != nil {
		t.Errorf("err:%v\n", err)
		return
	}
	reqest.Header.Add("content-type", "application/x-www-form-urlencoded")
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Encoding", "gzip, deflate")
	reqest.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

	r, err := c.Do(reqest)

	if err != nil {
		t.Errorf("err:%v\n", err)
		return
	}
	defer r.Body.Close()
	zlog.Debugf("rsp:%#v", r.Header)

	p, _ := ioutil.ReadAll(r.Body)
	fmt.Print(string(p))

}
