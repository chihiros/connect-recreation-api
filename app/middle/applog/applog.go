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
	fmt.Printf("logLevel: %v\n", logLevel)
	logrus.SetLevel(logLevel)
}

type AppLog struct{}

func NewLog() *AppLog {
	return &AppLog{}
}

func Print(args string) {
	pc, file, line, _ := runtime.Caller(2)
	fn := runtime.FuncForPC(pc).Name()
	msg := fmt.Sprintf("%s %s:%d %s", fn, file, line, args)
	log.Print(msg)
}

func (l *AppLog) Debug(args string) {
	fn, file, line := getCaller()
	msg := genMessage(fn, file, line, args)
	logrus.Debug(msg)
}

func (l *AppLog) Info(args string) {
	fn, file, line := getCaller()
	msg := genMessage(fn, file, line, args)
	logrus.Info(msg)
}

func (l *AppLog) Warn(args string) {
	fn, file, line := getCaller()
	msg := genMessage(fn, file, line, args)
	logrus.Warn(msg)
}

func (l *AppLog) Error(args string) {
	fn, file, line := getCaller()
	msg := genMessage(fn, file, line, args)
	logrus.Error(msg)
}

func (l *AppLog) Panic(err error) {
	fn, file, line := getCaller()
	msg := genMessage(fn, file, line, err)
	logrus.Panic(msg)
}

// 実行元のファイル名と行数を取得
func getCaller() (funcName, filepath string, line int) {
	pc, file, line, _ := runtime.Caller(2)
	fn := runtime.FuncForPC(pc).Name()

	return fn, file, line
}

func genMessage(fn, file string, line int, args interface{}) string {
	msg := fmt.Sprintf("%s %s:%d %v", fn, file, line, args)
	return msg
}
