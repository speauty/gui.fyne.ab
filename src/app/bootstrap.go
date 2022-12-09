package app

import (
	"gui.fyne.ab/src/app/page"
	"gui.fyne.ab/src/common/cfg"
	"gui.fyne.ab/src/core/gui"
)

type App struct {
	gui *gui.AppGui
}

func (ap *App) Init() *App {
	ap.gui = gui.Api()
	ap.gui.Init(cfg.Api().Gui)

	mainWindow := new(gui.AppWindow).NewWindow(gui.Api(), "主体窗口", true, true)
	mainWindow.RegisterPages(page.ApiABPage().Init(mainWindow))

	_ = ap.gui.RegisterWindows(mainWindow)

	return ap
}

func (ap *App) Run() {
	if err := ap.gui.Run(); err != nil {
		panic(err)
	}
}
