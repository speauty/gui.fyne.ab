// Package log 日志封装
package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"runtime"
	"sync"
	"time"
)

const (
	defaultLogFile         = "logs/app.log"        // 默认日志文件
	defaultTimestampFormat = "2006-01-02 15:04:05" // 默认时间格式
	strWindows             = "windows"
)

var (
	api  *Factory
	once sync.Once
)

func Api() *Factory {
	once.Do(func() {
		api = new(Factory)
	})
	return api
}

type Factory struct {
	sets sync.Map
}

func (factory *Factory) New(cfg *Cfg) *Log {
	if _, exist := factory.sets.Load(cfg.Key); !exist {
		if "" == cfg.LogFile {
			cfg.LogFile = defaultLogFile
		}
		tmpLog := new(Log)
		tmpLog.init(cfg)
		factory.sets.Store(cfg.Key, tmpLog)
		fmt.Println(fmt.Sprintf("[%s]日志服务[key: %s]注册成功", now(), cfg.Key))
	}
	currentDb, _ := factory.sets.Load(cfg.Key)
	return currentDb.(*Log)
}

func (factory *Factory) Get(key string) (*logrus.Logger, error) {
	if "" == key {
		key = "default"
	}
	if currentLog, exist := factory.sets.Load(key); exist {
		return currentLog.(*Log).client, nil
	}
	return nil, fmt.Errorf("[%s]当前日志服务[key: %s]暂未注册", now(), key)
}

type Log struct {
	client *logrus.Logger
	cfg    *Cfg
}

func (log *Log) AddHook(hook logrus.Hook) {
	log.client.Hooks.Add(hook)
}

func (log *Log) init(cfg *Cfg) {
	if nil == log.client {
		log.client = logrus.New()
		log.cfg = cfg
		log.set()
	}
}

func (log *Log) set() {
	if log.cfg.Level < 7 {
		log.client.SetLevel(logrus.Level(log.cfg.Level))
	}
	log.client.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: defaultTimestampFormat, DisableColors: false,
		ForceColors: true, FullTimestamp: true})
	if log.cfg.LogFile != "" { // 如果日志文件非空, 将日志打到对应文件
		var fd *rotatelogs.RotateLogs
		optLogFileFmt := log.cfg.LogFile + ".%Y%m%d"
		optWithLinkName := rotatelogs.WithLinkName(log.cfg.LogFile)
		optWithMaxAge := rotatelogs.WithMaxAge(time.Duration(log.cfg.LogMaxAge) * time.Second)
		optWithRotationCount := rotatelogs.WithRotationCount(log.cfg.LogRotationCount)
		if log.cfg.LogRotationTime == 0 {
			log.cfg.LogRotationTime = 60 * 60 * 24
		}
		optWithRotationTime := rotatelogs.WithRotationTime(time.Duration(log.cfg.LogRotationTime) * time.Second)
		var opts []rotatelogs.Option
		if runtime.GOOS != strWindows {
			opts = append(opts, optWithLinkName)
		}
		opts = append(opts, optWithRotationTime)
		if log.cfg.LogRotationCount > 0 {
			opts = append(opts, optWithRotationCount)
		} else {
			opts = append(opts, optWithMaxAge)
		}
		fd, _ = rotatelogs.New(optLogFileFmt, opts...)

		log.client.SetFormatter(&logrus.JSONFormatter{TimestampFormat: defaultTimestampFormat})
		log.client.SetOutput(fd)
	}
}

func now() string {
	return time.Now().Format(defaultTimestampFormat)
}
