package log

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logrus.Logger
}

var (
	log  *Logger
	once sync.Once
)

func Get() *Logger {
	once.Do(func() {
		log = &Logger{
			Logger: logrus.Logger{
				Out:       os.Stdout,
				Formatter: new(logrus.TextFormatter),
				Hooks:     make(logrus.LevelHooks),
				Level:     logrus.DebugLevel,
			},
		}
	})

	return log
}

var (
	Debug  = Get().Debug
	Debugf = Get().Debugf
	Info   = Get().Info
	Infof  = Get().Infof
	Warn   = Get().Warn
	Warnf  = Get().Warnf
	Error  = Get().Error
	Errorf = Get().Errorf
	Fatal  = Get().Fatal
	Fatalf = Get().Fatalf
	Panic  = Get().Panic
	Panicf = Get().Panicf
)
