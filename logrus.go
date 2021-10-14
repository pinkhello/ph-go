package phgo

import (
	"fmt"
	"time"

	rotates "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const (
	TimestampFormat = "2006-01-02 15:04:05"
)

var (
	appName = GetEnv("APP_NAME", "-")
)

func SetAppName(name string) {
	appName = name
}

type (
	AppNameFieldHook struct {
	}

	Logger struct {
		Log     *logrus.Logger
		LogType string
		LogPath string
	}
)

func (hook *AppNameFieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["appName"] = appName
	return nil
}

func (hook *AppNameFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *Logger) Init(fields logrus.Fields) {
	logger := logrus.New()
	logger.SetReportCaller(true)
	formatter := &logrus.JSONFormatter{
		TimestampFormat: TimestampFormat,
	}
	logger.SetFormatter(formatter)
	if fields != nil {
		logger.WithFields(fields)
	}
	logger.AddHook(&AppNameFieldHook{})
	logger.AddHook(NewLfsHook(fmt.Sprintf("%s/%s.log", l.LogPath, l.LogType), formatter))
	l.Log = logger
}

func NewLfsHook(path string, formatter logrus.Formatter) *lfshook.LfsHook {
	writer, err := rotates.New(
		path+".%Y%m%d",
		rotates.WithLinkName(path),
		rotates.WithMaxAge(time.Duration(15*24)*time.Hour),
		rotates.WithRotationTime(time.Duration(24)*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		}, formatter)
}
