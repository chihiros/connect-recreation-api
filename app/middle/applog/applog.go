package applog

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

func SetEnv(env ENV) {
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
	msg := genMessage(args)
	logrus.Debug(msg)
}

func Info(args string) {
	msg := genMessage(args)
	logrus.Info(msg)
}

func Warn(args string) {
	msg := genMessage(args)
	logrus.Warn(msg)
}

func Error(args string) {
	msg := genMessage(args)
	logrus.Error(msg)
}

func Panic(err error) {
	msg := genMessage(err)
	logrus.Panic(msg)
}

// 実行元のファイル名と行数を取得
func getCaller() (funcName, filepath string, line int) {
	pc, file, line, _ := runtime.Caller(3)
	fn := runtime.FuncForPC(pc).Name()

	return fn, file, line
}

func genMessage(args interface{}) string {
	fn, file, line := getCaller()
	msg := fmt.Sprintf("%s %s:%d %v", fn, file, line, args)
	return msg
}
