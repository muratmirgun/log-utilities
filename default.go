package log_utilities

import (
	"github.com/sirupsen/logrus"
	"time"
)

func Default(f *Formatter) error {
	f.SeverityMap = map[string]string{
		"panic":   logrus.PanicLevel.String(),
		"fatal":   logrus.FatalLevel.String(),
		"warning": logrus.WarnLevel.String(),
		"debug":   logrus.DebugLevel.String(),
		"error":   logrus.ErrorLevel.String(),
		"trace":   logrus.TraceLevel.String(),
		"info":    logrus.InfoLevel.String(),
	}

	f.TimestampFormat = func(fields logrus.Fields, now time.Time) error {
		ts := time.Now().Format(time.StampMicro)
		fields["timestamp"] = ts
		return nil
	}
	return nil
}
