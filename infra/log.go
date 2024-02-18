package infra

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

var logger *logrus.Logger

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func NewLogger() *logrus.Logger {
	if logger == nil {
		path := "log/"

		isExist, err := dirExists(path)
		if err != nil {
			panic(err)
		}

		if !isExist {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		writer, err := rotatelogs.New(
			path+"AppNameGo"+"-"+"%Y%m%d.log",
			rotatelogs.WithMaxAge(-1),
			rotatelogs.WithRotationCount(uint(time.Duration(24)*time.Hour)),
			rotatelogs.WithRotationTime(4),
		)
		if err != nil {
			panic(err)
		}

		logger = logrus.New()

		// TODO: Active this code if later it's needed to limit the log
		// // Set Log level that need to show or stored
		// if conf.App.Environtment == constants.EnvProd {
		// 	logger.SetLevel(logrus.WarnLevel)
		// } else {
		// 	logger.SetLevel(logrus.DebugLevel)
		// }

		// Set Hook with writer & formatter for log file
		logger.Hooks.Add(lfshook.NewHook(
			writer,
			&logrus.TextFormatter{
				DisableColors:   false,
				FullTimestamp:   true,
				TimestampFormat: "2006-01-02 15:04:05",
			},
		))

		// Set formatter for os.Stdout
		logger.SetFormatter(&logrus.TextFormatter{
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})

		logger.SetFormatter(&ecslogrus.Formatter{})

		return logger
	}

	return logger
}

func TestNewLogger() *logrus.Logger {
	logger := logrus.New()
	return logger
}
