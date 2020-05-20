package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go_schedule/consts"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"
)

func InitLogger() {
	beginTime := time.Now()
	logInit(&beginTime)
	logrus.Infof("Task begin at: %v", beginTime.Format(consts.TimeStr))
}

//logFile 日志文件对象
var logFile *os.File

//logInit 初始化日志组件
func logInit(logTime *time.Time) {
	logFilePath, _ := filepath.Abs(consts.LogFilePath)
	logFileName := consts.LogFilePrefix + logTime.Format(consts.LogFileTimeStr) + consts.LogFileSuffix
	fileName := path.Join(logFilePath, logFileName)
	logFile, _ = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)

	logrus.SetOutput(io.MultiWriter(os.Stderr, logFile))
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: consts.TimeStr,
	})

	//exit方法禁止使用
	//logrus.RegisterExitHandler(closeLogFile)
}

//CloseLogFile 关闭日志文件对象
func CloseLogFile() {
	logrus.Infof("Task end at: %v", time.Now().Format(consts.TimeStr))
	//美化控制台输出
	fmt.Println()

	if logFile != nil {
		logFile.Close()
	}
}
