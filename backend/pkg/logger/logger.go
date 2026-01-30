package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	log  *logrus.Logger
	once sync.Once
)

func New() *logrus.Logger {
	once.Do(func() {
		log = logrus.New()
		log.SetLevel(logrus.DebugLevel)
		log.SetFormatter(&logrus.JSONFormatter{})
	})

	return log
}
