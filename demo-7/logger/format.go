package logger

import (
	"github.com/sirupsen/logrus"
	requestid "grpc-demo/demo-7/requestId"
)

func init() {
	logrus.SetFormatter(&logFormatter{
		JSONFormatter: &logrus.JSONFormatter{},
	})
}

type logFormatter struct {
	*logrus.JSONFormatter
}

func (f *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	requestId := requestid.RequestID(entry.Context)

	if len(requestId) > 0 {
		entry.Data["x-request-id"] = requestId
	}

	return f.JSONFormatter.Format(entry)
}
