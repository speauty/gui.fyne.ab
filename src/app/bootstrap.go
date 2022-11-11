package app

import (
	"gui.fyne.ab/src/common/cfg"
	"gui.fyne.ab/src/core/gui"
)

var (
	MainWindow *gui.AppWindow
)

type App struct {
	gui *gui.AppGui
}

func (ap *App) Init() *App {
	ap.gui = gui.Api()
	ap.gui.Init(cfg.Api().Gui)

	ap.initWindows()

	_ = ap.gui.RegisterWindows(genWindow())

	return ap
}

func (ap *App) Run() {
	if err := ap.gui.Run(); err != nil {
		panic(err)
	}
}

func (ap *App) initWindows() {
	MainWindow = new(gui.AppWindow).NewWindow(gui.Api(), "主体窗口", true, true)
}
