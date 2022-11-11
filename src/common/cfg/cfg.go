package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"gui.fyne.ab/src/core/gui"
	"gui.fyne.ab/src/core/log"
	"sync"
)

var (
	api  *Cfg
	once sync.Once
)

func Api() *Cfg {
	once.Do(func() {
		api = new(Cfg)
	})
	return api
}

type Cfg struct {
	App struct {
	} `json:"app" yaml:"app" xml:"app" mapstructure:"app"`
	Gui *gui.Cfg `json:"gui" yaml:"gui" xml:"gui" mapstructure:"gui"`
	Log *log.Cfg `json:"log" yaml:"log" xml:"log" mapstructure:"log"`
}

func (cfg *Cfg) LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("当前配置[路径: %s]载入失败, 错误: %s", path, err.Error()))
	}

	if err := viper.Unmarshal(cfg); err != nil {
		panic(fmt.Errorf("默认配置[路径: %s]JSON解析失败, 错误: %s", path, err.Error()))
	}

	return
}

func (cfg *Cfg) GetAppName() string {
	return cfg.Gui.AppName
}
