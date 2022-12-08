package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gui.fyne.ab/src/common/snowman"
	"gui.fyne.ab/src/core/gui"
	"sync"
)

var (
	envApi  *EnvPage
	envOnce sync.Once
)

func ApiEnvPage() *EnvPage {
	envOnce.Do(func() {
		envApi = new(EnvPage)
	})
	return envApi
}

type EnvPage struct {
	BasePage
}

func (tp *EnvPage) Init(window *gui.AppWindow) gui.IPage {
	tp.window = window
	tp.id = snowman.NewSnowApi().GetIdInt64()
	tp.name = "环境"
	tp.flagIsError = false
	tp.flagIsStart = false

	return tp
}

func (tp *EnvPage) GenCanvasObject() fyne.CanvasObject {
	return container.New(layout.NewCenterLayout(), widget.NewLabel("环境页面"))
}
