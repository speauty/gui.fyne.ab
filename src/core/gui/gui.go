package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"sync"
)

var (
	api  *AppGui
	once sync.Once
)

func Api() *AppGui {
	once.Do(func() {
		defer func() {
			fmt.Println(fmt.Sprintf("[%s]初始化GUI应用: 成功", now()))
		}()
		api = new(AppGui)
	})
	return api
}

// AppGui 应用界面 单例模式, 一个软件, 只支持一个应用, 所以直接这里就限制
type AppGui struct {
	app        fyne.App   // fyne应用
	cfg        *Cfg       // 配置
	state      *state     // 状态
	windows    sync.Map   // 注册窗口集合
	mainWindow *AppWindow // 主窗口
}

// GetCfg 获取当前应用配置
func (ag *AppGui) GetCfg() *Cfg { return ag.cfg }

// GetAppName 获取当前应用名称
func (ag *AppGui) GetAppName() string { return ag.cfg.AppName }

// RegisterWindows 注册窗口, 不采用循环协程处理
func (ag *AppGui) RegisterWindows(windows ...*AppWindow) error {
	for _, win := range windows {

		if _, isExisted := ag.windows.Load(win.GetWindowId()); isExisted { // 已注册, 判断之后, 直接中断
			return fmt.Errorf(
				"[%s]窗口[Id: %d, Name: %s]已在应用[%s]注册, 请勿重复注册", now(),
				win.GetWindowId(), win.GetWindowName(), ag.GetAppName(),
			)
		}

		ag.windows.Store(win.GetWindowId(), win)

		if win.IsMainWindow() { // 主窗口注册处理
			if nil == ag.mainWindow {
				ag.mainWindow = win
				fmt.Println(fmt.Sprintf(
					"[%s]应用[%s]注册主窗口[Id: %d, Name: %s]: 成功", now(),
					ag.GetAppName(), win.GetWindowId(), win.GetWindowName(),
				))
			} else {
				return fmt.Errorf(
					"[%s]应用[%s]已注册主窗口[Id: %d, Name: %s], 当前尝试注册新主窗口[Id: %d, Name: %s]", now(),
					ag.GetAppName(), ag.mainWindow.GetWindowId(), ag.mainWindow.GetWindowName(),
					win.GetWindowId(), win.GetWindowName(),
				)
			}
		}

		fmt.Println(fmt.Sprintf(
			"[%s]窗口[Id: %d, Name: %s]在应用[%s]注册: 成功", now(),
			win.GetWindowId(), win.GetWindowName(), ag.GetAppName(),
		))
	}
	return nil
}

func (ag *AppGui) Init(cfg *Cfg) *AppGui {
	if nil == ag.app { // 创建fyne应用
		ag.app = app.New()
	}
	if nil == ag.state { // 生成默认状态
		ag.state = new(state).GenDefault()
	}

	if nil == ag.cfg {
		ag.cfg = cfg
	}

	go ag.setIcon()

	return ag
}

func (ag *AppGui) Run() error {
	if nil == ag.mainWindow {
		return fmt.Errorf("[%s]当前应用[%s]暂未注册主窗口, 启动失败", now(), ag.GetAppName())
	}
	ag.mainWindow.OnRender()
	return nil
}

func (ag *AppGui) setIcon() {
	if "" != ag.cfg.AppIconPath {
		iconResource, err := fyne.LoadResourceFromPath(ag.cfg.AppIconPath)
		if nil != err {
			fmt.Println(fmt.Errorf("[%s]当前应用[%s]注册图标资源[%s]失败", now(), ag.GetAppName(), ag.cfg.AppIconPath))
			return
		}
		ag.app.SetIcon(iconResource)
	}
}
