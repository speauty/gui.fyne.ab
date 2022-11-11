package gui

// 窗口

import (
	"fmt"
	"fyne.io/fyne/v2"
	"math/rand"
	"sync"
)

type AppWindow struct {
	gui              *AppGui
	window           fyne.Window
	windowId         int64  // 窗口主键
	windowName       string // 窗口名称
	flagIsMainWindow bool   // 是否是主窗口
	flagFixedSize    bool   // 是否固定尺寸
	flagIsRendering  bool   // 是否正在渲染
	currentPageId    int64
	pages            sync.Map
}

// NewWindow 创建新窗口
func (aw *AppWindow) NewWindow(gui *AppGui, windowName string, isMain bool, isFixedSize bool) *AppWindow {
	aw.windowId = rand.Int63()
	aw.gui = gui
	aw.windowName = windowName
	aw.flagIsMainWindow = isMain
	aw.flagFixedSize = isFixedSize
	aw.flagIsRendering = false
	aw.currentPageId = 0
	aw.setWindow()
	return aw
}

func (aw *AppWindow) GetWindowId() int64 { return aw.windowId }

func (aw *AppWindow) GetWindowName() string { return aw.windowName }

func (aw *AppWindow) GetWindow() fyne.Window { return aw.window }

func (aw *AppWindow) IsMainWindow() bool { return aw.flagIsMainWindow }
func (aw *AppWindow) OpStopRender()      { aw.flagIsRendering = false }
func (aw *AppWindow) OpStartRender()     { aw.flagIsRendering = true }

func (aw *AppWindow) OnRender() {
	go func() {
		for true {
			aw.OnUpdate() // 更新

			if aw.flagIsRendering { // swap-buffers
				aw.window.Content().Refresh()
				aw.flagIsRendering = false
				aw.currentPageId = 0
			}
		}
	}()
	if aw.flagIsMainWindow { // 主窗口的话, 需要使用ShowAndRun, 挂起个主循环
		aw.window.ShowAndRun()
	} else {
		aw.window.Show()
	}
}

func (aw *AppWindow) OnUpdate() {
	aw.pages.Range(func(pageId, page any) bool {
		if pageId == aw.currentPageId {
			aw.SetContent(page.(IPage).GenCanvasObject())
			aw.flagIsRendering = true
			fmt.Println(fmt.Sprintf("[%s]窗口[%s]准备载入页面[%s]", now(), aw.GetWindowName(), page.(IPage).GetName()))
			return false
		}
		return true
	})
}

func (aw *AppWindow) RegisterPages(pages ...IPage) {
	for _, page := range pages { // @todo 不判断重复插入之类的
		aw.pages.Store(page.GetId(), page)
		if page.FlagIsStartPage() {
			aw.SetCurrentPageId(page.GetId())
		}
		fmt.Println(fmt.Sprintf(
			"[%s]页面[Name: %s]在窗口[%s]注册: 成功", now(), page.GetName(), aw.GetWindowName(),
		))
	}
}

// SetContent 设置内容
func (aw *AppWindow) SetContent(canvasObj fyne.CanvasObject) {
	aw.window.SetContent(canvasObj)
}

// SetMainMenu 设置主菜单
func (aw *AppWindow) SetMainMenu(mainMenu *fyne.MainMenu) {
	aw.window.SetMainMenu(mainMenu)
}

// SetCurrentPageId 设置加载页面ID
func (aw *AppWindow) SetCurrentPageId(pageId int64) {
	aw.currentPageId = pageId
}

// SetCloseFunc 设置关闭处理句柄
func (aw *AppWindow) SetCloseFunc(fn func()) {
	if nil != fn {
		aw.window.SetCloseIntercept(fn)
	}
}

func (aw *AppWindow) setWindow() {
	if nil == aw.window {
		app := aw.gui.app
		if aw.flagIsMainWindow {
			aw.windowName = aw.gui.GetCfg().AppName
		}
		aw.window = app.NewWindow(aw.windowName)
		aw.window.Resize(fyne.NewSize(800, 600))
		aw.window.SetFixedSize(aw.flagFixedSize) // 设置是否固定尺寸, 按理说应该放在主窗口判断内的, 但为了避免窗口子窗口之类的情况
		if aw.flagIsMainWindow {                 // 主窗口配置
			aw.window.SetMaster()
		}
	}
}
