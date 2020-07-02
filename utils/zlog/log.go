package zlog

import (
	"github.com/astaxie/beego/validation"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Debugr 不加sugar的方法，可添加字段
func Debugr(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

//Debugf 加sugar的方法
func Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}

//Debugw 加sugar的方法
func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Debugw(msg, keysAndValues...)

}

//Debug 加sugar的方法
func Debug(args ...interface{}) {
	logger.Sugar().Debug(args...)
}

//Infor 不加sugar的方法，可添加字段
func Infor(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

//Infof 加sugar的方法
func Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

//Infow 加sugar的方法
func Infow(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Infow(msg, keysAndValues...)

}

//Info 加sugar的方法
func Info(args ...interface{}) {
	logger.Sugar().Info(args...)
}

//Warnr 不加sugar的方法，可添加字段
func Warnr(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

//Warnf 加sugar的方法
func Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}

//Warnw 加sugar的方法
func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Warnw(msg, keysAndValues...)

}

//Warn 加sugar的方法
func Warn(args ...interface{}) {
	logger.Sugar().Warn(args...)
}

//Errorr 不加sugar的方法，可添加字段
func Errorr(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

//Errorf 加sugar的方法
func Errorf(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

//Errorw 加sugar的方法
func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Errorw(msg, keysAndValues...)

}

//Error 加sugar的方法
func Error(args ...interface{}) {
	logger.Sugar().Error(args...)
}

//DPanicr 不加sugar的方法，可添加字段
func DPanicr(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

//DPanicf 加sugar的方法
func DPanicf(template string, args ...interface{}) {
	logger.Sugar().DPanicf(template, args...)
}

//DPanicw 加sugar的方法
func DPanicw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().DPanicw(msg, keysAndValues...)

}

//DPanic 加sugar的方法
func DPanic(args ...interface{}) {
	logger.Sugar().DPanic(args...)
}

//Panicr 不加sugar的方法，可添加字段
func Panicr(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

//Panicf 加sugar的方法
func Panicf(template string, args ...interface{}) {
	logger.Sugar().Panicf(template, args...)
}

//Panicw 加sugar的方法
func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Panicw(msg, keysAndValues...)

}

//Panic 加sugar的方法
func Panic(args ...interface{}) {
	logger.Sugar().Panic(args...)
}

//Fatalr 不加sugar的方法，可添加字段
func Fatalr(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

//Fatalf 加sugar的方法
func Fatalf(template string, args ...interface{}) {
	logger.Sugar().Fatalf(template, args...)
}

//Fatalw 加sugar的方法
func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Fatalw(msg, keysAndValues...)

}

//Fatal 加sugar的方法
func Fatal(args ...interface{}) {
	logger.Sugar().Fatal(args...)
}

//Sync 执行logger.Sync(),程序结束时，将缓存中数据都打印出来。
func Sync() error {
	return logger.Sync()
}

// With creates a child logger and adds structured context to it. Fields added
func With(fields ...zap.Field) *zap.Logger {
	return logger.With(fields...)
}

// Check 调用logger的Check,如果日志级别需要打印，则返回检查过的入口
func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return logger.Check(lvl, msg)
}

// Core returns 返回logger下的 zapcore.Core.
func Core() zapcore.Core {
	return logger.Core()
}

//MutiErrors 多个错误输出，用于beego的表单验证包
func MutiErrors(errors []*validation.Error) {

	for _, err := range errors {

		logger.Sugar().Error(err.Key, err.Message)

	}

}
