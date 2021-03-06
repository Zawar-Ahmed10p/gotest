package gologger

import (
	"fmt"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

//Logger tem

var Logger *logrus.Logger

//GetLogger returns logger
func GetNewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "err.log",
		MaxSize:    1, //file size in megabytes
		MaxBackups: 5, //max number of files
		MaxAge:     2, //days
	})
	logger.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%time% [%lvl%]: %msg% \n",
	})
	Logger = logger
	return Logger

}

var initOnce sync.Once

func GetLogger() {
	initOnce.Do(func() {
		fmt.Println("Logger initialized")
		logger := logrus.New()
		logger.SetOutput(&lumberjack.Logger{
			Filename:   "err.log",
			MaxSize:    1, //file size in megabytes
			MaxBackups: 5, //max number of files
			MaxAge:     2, //days
		})
		logger.SetFormatter(&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% [%lvl%]: %msg% \n",
		})
		Logger = logger

	})

}
