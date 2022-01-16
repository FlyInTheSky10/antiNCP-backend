package log

import (
	"antiNCP/config"
	"fmt"
	rotateLogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var Logger *logrus.Logger

func init() {
	Logger = getLogger()
	if Logger == nil {
		panic("Logger failed to initialize.")
	}
	Logger.Info("Logger successfully initialized.")
}

func getLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	//记录调用位置：调用文件+调用行号+调用函数
	//默认开启，开启会有20%-40%的性能损失
	//如不想开启还想记录报错位置，可以对[Logger]进行进一步封装，见文件末尾
	logger.SetReportCaller(true)

	if config.C.Debug == true {
		logger.SetLevel(logrus.DebugLevel)
	}

	logConfig := config.C.LogConf
	if _, err := os.Stat(logConfig.LogPath); os.IsNotExist(err) {
		err := os.MkdirAll(logConfig.LogPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	baseLogPath := path.Join(logConfig.LogPath, logConfig.LogFileName)
	writer, err := rotateLogs.New(
		baseLogPath+".%Y-%m-%d-%H-%M",
		rotateLogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotateLogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotateLogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		logger.Fatal(err)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})

	logger.AddHook(lfHook)
	return logger
}
