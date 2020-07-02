package zlog

import (
	"os"
	"sz_resume_202005/utils/setting"

	"github.com/siddontang/go/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//Logger logger实例
var logger *zap.Logger

//Init 初始化zap日志库，参看zap package-level example advancedConfigurion
func Init() {

	var (
		src        string
		fileOutput zapcore.WriteSyncer
	)
	if setting.EnvRoot != "" {
		src = setting.EnvRoot + getLogFilePath() + getLogFileName()
		fileOutput = zapcore.Lock(zapcore.AddSync(&lumberjack.Logger{
			Filename:   src,
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
		}))
	} else {
		fileOutput = zapcore.Lock(zapcore.AddSync(&lumberjack.Logger{
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
		}))
		log.Errorf("envPath not exsit,can't find logPath\n")

	}

	consoleOutput := zapcore.AddSync(os.Stdout)

	//高优先级
	// highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	// 	return lvl >= zapcore.DebugLevel
	// })
	// //低优先级
	// lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	// 	return lvl < zapcore.ErrorLevel
	// })

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.

	fileEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		//EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileOutput, mapToLevel(setting.AppSetting.LogLevel)),
		zapcore.NewCore(consoleEncoder, consoleOutput, zap.DebugLevel),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func mapToLevel(s string) (l zapcore.LevelEnabler) {
	switch s {

	case "INFO":
		l = zap.InfoLevel
	case "WARN":
		l = zap.WarnLevel
	case "ERROR":
		l = zap.ErrorLevel
	case "DPANIC":
		l = zap.DPanicLevel
	case "PANIC":
		l = zap.PanicLevel
	case "FATAL":
		l = zap.FatalLevel
	default:
		l = zap.DebugLevel

	}
	return
}
