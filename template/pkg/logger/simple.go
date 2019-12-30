package logger

import (
	"os"

	"github.com/op/go-logging"
)

type simpleLogger struct {
	log *logging.Logger
}

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// NewSimpleLogger initializes the simple version logger by go-logging lib
func NewSimpleLogger() Log {
	profile := logging.NewLogBackend(os.Stderr, "", 0)
	profileFormatter := logging.NewBackendFormatter(profile, format)
	logging.SetBackend(profileFormatter)
	var log = logging.MustGetLogger("")
	return &simpleLogger{
		log: log,
	}
}

func (l *simpleLogger) Log(vals ...interface{}) error {
	l.log.Info(vals)
	return nil
}

func (l *simpleLogger) Debug(vals ...interface{}) error {
	l.log.Debug(vals)
	return nil
}

func (l *simpleLogger) Info(vals ...interface{}) error {
	l.log.Info(vals)
	return nil
}

func (l *simpleLogger) Warn(vals ...interface{}) error {
	l.log.Warning(vals)
	return nil
}

func (l *simpleLogger) Error(vals ...interface{}) error {
	l.log.Error(vals)
	return nil
}

func (l *simpleLogger) Errorf(str string, vals ...interface{}) error {
	l.log.Errorf(str, vals...)
	return nil
}
