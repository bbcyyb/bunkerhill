package simple

import (
	"github.com/bbcyyb/bunkerhill/logs"
)

var logger = logs.NewLogger()

func GetSimpleLogger() *logs.LoggerWrapper {
	return logger
}

func Async(msgLen ...int64) *logs.LoggerWrapper {
	return logger.Async(msgLen...)
}

func SetLevel(l int) {
	logger.SetLevel(l)
}

func SetLogger(adapter string, config ...string) error {
	return logger.SetLogger(adapter, config...)
}

func Critical(f interface{}, v ...interface{}) {
	logger.Critical(logs.FormatLog(f, v...))
}

func Error(f interface{}, v ...interface{}) {
	logger.Error(logs.FormatLog(f, v...))
}

func Wran(f interface{}, v ...interface{}) {
	logger.Wran(logs.FormatLog(f, v...))
}

func Info(f interface{}, v ...interface{}) {
	logger.Info(logs.FormatLog(f, v...))
}

func Debug(f interface{}, v ...interface{}) {
	logger.Debug(logs.FormatLog(f, v...))
}
