package main

import (
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	//"sz_resume_202005/dao/mysql"
	"sz_resume_202005/dao/db"
	"sz_resume_202005/dao/redisdb"
	"sz_resume_202005/routers"
	"sz_resume_202005/utils/middleware"
	"sz_resume_202005/utils/setting"
	"sz_resume_202005/utils/upload"
	"sz_resume_202005/utils/zlog"

	"github.com/gin-gonic/gin"
)

func main() {
	setting.Init()
	zlog.Init()
	defer zlog.Sync()
	redisdb.Init()
	db.Init()

	gin.SetMode(setting.RunMode) //默认的模式就是debug，但是为了好习惯，还是要加上。

	r := gin.Default()
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.Use(middleware.Cors(), middleware.RequestInfo()) //跨域访问中间件
	routers.LoadRouters(r)
	if err := r.Run(setting.Server); err != nil {
		zlog.Errorf("startup service failed, err:%v", err)
	}

}
func gtAppPath() string {

	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
