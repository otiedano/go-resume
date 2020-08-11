package setting

import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

var (
	//EnvRoot 系统环境变量，根目录路径
	EnvRoot string

	//EnvRootName 项目必须指定环境变量方可正常运行，日志库依赖于此变量，读取配置文件也一定程度依赖此变量。
	EnvRootName = "SZRESUME"

	// Cfg 配置文件
	Cfg *ini.File
	//MysqlAddr mysql地址
	MysqlAddr string
	// RedisAddr redis地址
	RedisAddr string
	//RunMode 运行模式
	RunMode string
	//PageSize 页码
	PageSize int
	//Server server地址或端口
	Server string
	//ReadTimeout 读超时
	ReadTimeout time.Duration
	//WriteTimeout 写超时
	WriteTimeout time.Duration
)

//AppSetting app设置
var AppSetting = &App{}

//App app设置
type App struct {
	Domain          string
	PrefixURL       string
	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	LogLevel    string
	TimeFormat  string

	TokenExpire int
}

var (
	mysqlAddr = "root:tiedano1@tcp(127.0.0.1:3306)/resume?charset=utf8mb4&loc=Asia%2FShanghai&parseTime=true"
	redisAddr = "localhost:6379"

	cfgApp = "conf/app_dev.ini"
)

//Init 初始化，init的情况test不通过
func Init() {
	var err error
	EnvRoot = os.Getenv(EnvRootName) //配置文件和日志库多处引用，所以需要先执行。
	// log.Printf("EnvRoot:%v", os.Getenv("HOME"))
	// log.Printf("HOME:%v", EnvRoot)
	// log.Printf("读取cfg内容失败，sz_golang:%v", err)
	Cfg, err = ini.Load(cfgApp)
	if err != nil {
		log.Printf("读取cfg内容失败，sz_golang:%v", err)
		//log.Printf("Fail to parse 'conf/app.ini': %v", err)
		if EnvRoot != "" {
			src := EnvRoot + cfgApp
			Cfg, err = ini.Load(src)
			if err != nil {
				log.Fatalf("Fail to parse '%v' with absolute path: %v", cfgApp, err)
			}
		} else {
			log.Fatalf("Fail to parse '%v',can not read envpath.", cfgApp)
		}

	} else {
		log.Printf("读取cfg内容失败，sz_golang:%v", err)
	}
	loadDefault()

	loadMysql()
	loadRedis()
	loadServer()
	mapTo("app", AppSetting)

}
func loadMysql() {
	sec, err := Cfg.GetSection("mysql")
	if err != nil {
		log.Fatalf("Fail to get section 'mysql': %v", err)
	}

	MysqlAddr = sec.Key("address").MustString(mysqlAddr)
}
func loadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get section 'redis': %v", err)
	}

	RedisAddr = sec.Key("address").MustString(redisAddr)
}
func loadDefault() {
	RunMode = Cfg.Section("").Key("run_mode").MustString("debug")
	PageSize = Cfg.Section("").Key("page_size").MustInt(10)
}
func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	Server = ":" + sec.Key("port").MustString("8080")
	ReadTimeout = time.Duration(sec.Key("read_timeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("write_timeout").MustInt(60)) * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := Cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
