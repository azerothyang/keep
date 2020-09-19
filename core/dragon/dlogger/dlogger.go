package dlogger

import (
	"dragon/core/dragon/conf"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

// write log
func writeLog(data interface{}, level string) {
	if data == nil {
		// 如果data为空，不进行打印
		return
	}
	// 根据data类型删除json或者字符串
	now := time.Now()
	datetime := now.Format("2006-01-02 15:04:05")
	date := now.Format("2006-01-02")
	// 生成或打开文件
	logDir := conf.Conf.Log.Dir
	path := conf.ExecDir + "/" + logDir + "/" + date + "." + level + ".log"
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	var logInfo string
	if reflect.TypeOf(data).String() == "string" {
		logInfo = data.(string)
	} else {
		d, _ := json.Marshal(data)
		logInfo = string(d)
	}
	// todo check if safe
	go func(file *os.File, datetime string, level string, logInfo string) {
		fmt.Fprintf(file, "[%s] [%s] || %s \r\n\r\n", datetime, level, logInfo)
		file.Close()
	}(logFile, datetime, level, logInfo)
}

func Debug(data interface{}) {
	writeLog(data, "debug")
}

func Info(data interface{}) {
	writeLog(data, "info")
}

func Warn(data interface{}) {
	writeLog(data, "warn")
}

func Error(data interface{}) {
	writeLog(data, "error")
}

func SqlInfo(data interface{}) {
	writeLog(data, "sql")
}

func SqlError(data interface{}) {
	writeLog(data, "sql.error")
}
