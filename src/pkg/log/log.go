package log

import (
	"github.com/hamidteimouri/gommon/htcolog"
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	level, err := logrus.ParseLevel(htenvier.Env("LOG_LEVEL"))
	if err != nil {
		e := htcolog.MakeRed("LOG_LEVEL is not valid")
		panic(e)
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(customLogger{
		service: htenvier.Env("APP_NAME"),
		formatter: &logrus.JSONFormatter{
			PrettyPrint:     true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	})

	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
}

type customLogger struct {
	service   string
	formatter logrus.Formatter
}

func (l customLogger) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Data["service"] = l.service
	return l.formatter.Format(entry)
}
