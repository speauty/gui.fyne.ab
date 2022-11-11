package app

import (
	"github.com/sirupsen/logrus"
	"gui.fyne.ab/src/common/cfg"
	"gui.fyne.ab/src/core/log"
	"sync"
)

var (
	logAppApi  *logrus.Logger
	logAppOnce sync.Once
)

func LogAppApi() *logrus.Logger {
	logAppOnce.Do(func() {
		logAppApi = new(logrus.Logger)
		_ = log.Api().New(cfg.Api().Log)
		logAppApi, _ = log.Api().Get(cfg.Api().Log.Key)

	})
	return logAppApi
}
