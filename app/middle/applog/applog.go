package applog

/*
* Use Example
applog.Debug("Server started at :8080")
applog.Info("Server started at :8080")
applog.Warn("Server started at :8080")
applog.Error("Server started at :8080")
*/

import (
	"fmt"
	"runtime"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var logLevel logrus.Level

type ENV int

const (
	PROD ENV = iota
	STG
	DEV
)

func Setenv(env ENV) {
	switch env {
	case PROD:
		logLevel = logrus.WarnLevel
	case STG:
		logLevel = logrus.TraceLevel
	case DEV:
		logLevel = logrus.TraceLevel
	default:
		logLevel = logrus.TraceLevel
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	// logrus.SetFormatter(&logrus.JSONFormatter{}) // 出力の形式がJSONになる
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetLevel(logLevel)
}

func Debug(args string) {
	logrus.Debug(args)
}

func Info(args string) {
	logrus.Info(genMessage(args))
}

func Warn(args string) {
	logrus.Warn(genMessage(args))
}

func Error(args string) {
	logrus.Error(genMessage(args))
}

func Panic(err error) {
	logrus.Panic(genMessage(err))
}

// 実行元のファイル名と行数を取得
func getCaller() (funcName, filepath string, line int) {
	pc, file, line, _ := runtime.Caller(3)
	fn := runtime.FuncForPC(pc).Name()
	return fn, file, line
}

func genMessage(args interface{}) string {
	fn, file, line := getCaller()
	msg := fmt.Sprintf("%s:%d %s %v", file, line, fn, args)
	return msg
}
