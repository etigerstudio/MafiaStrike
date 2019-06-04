package util

import (
	"fmt"
	"mafia-strike/consts"
	"log"
)

func Infoln(v ...interface{}) {
	logln(consts.LogInfo, v...)
}

func Infof(format string, v ...interface{}) {
	logf(consts.LogInfo, format, v...)
}

func Warnln(v ...interface{}) {
	logln(consts.LogWarning, v...)
}

func Warnf(format string, v ...interface{}) {
	logf(consts.LogWarning, format, v...)
}

func Errorln(v ...interface{}) {
	logln(consts.LogError, v...)
}

func Errorf(format string, v ...interface{}) {
	logf(consts.LogError, format, v...)
}

func logln(category string, v ...interface{}) {
	log.Print(consts.LogPrefix + " " + category + " ", fmt.Sprintln(v...))
}

func logf(category string, format string, v ...interface{}) {
	log.Printf("%s %s", consts.LogPrefix + " " + category, fmt.Sprintf(format, v...))
}