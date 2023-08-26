package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Entry

type LogConfig struct {
	Path    string
	NameApp string
	isDebug bool
}

func InitLog(config LogConfig) *logrus.Entry {
	f, err := os.OpenFile(config.Path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic("failed to assign log file")
	}

	// don't forget to close it
	// defer f.Close()
	out := io.MultiWriter(os.Stdout, f)
	format := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "@caller",
			logrus.FieldKeyFile:  "@file",
		},
	}

	if config.isDebug {
		format.PrettyPrint = true
	}

	logger := logrus.Logger{
		Level:        logrus.InfoLevel,
		Out:          out,
		Formatter:    format,
		ReportCaller: true,
	}

	Logger = logger.WithFields(logrus.Fields{
		"@app": config.NameApp,
	})

	return Logger
}

func (l *LogConfig) SetDebug(isDebug bool) {
	l.isDebug = isDebug
}
