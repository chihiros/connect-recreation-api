package applog

import (
	"fmt"
	"runtime"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var logLevel string

const (
	PROD = iota
	STG
	DEV
)

	switch env {
	case "production":
		logLevel = "WARN"
	case "staging":
		logLevel = "DEBUG"
	default:
		logLevel = "DEBUG"
	}
}

type AppLog struct{}

func NewLog() *AppLog {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(logLevel),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	return &AppLog{}
}

func Print(args string) {
	pc, file, line, _ := runtime.Caller(2)
	fn := runtime.FuncForPC(pc).Name()
	msg := fmt.Sprintf("%s %s:%d %s", fn, file, line, args)
	log.Print(msg)
}

func (l *AppLog) Debug(args string) {
	Print("[DEBUG] " + args)
}

func (l *AppLog) Warn(args string) {
	Print("[WARN] " + args)
}

func (l *AppLog) Error(args string) {
	Print("[ERROR] " + args)
}

func (l *AppLog) Panic(err error) {
	Print("[PANIC] " + err.Error())
	panic(err)
}
