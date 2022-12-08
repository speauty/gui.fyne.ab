package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"gui.fyne.ab/src/common/snowman"
	"gui.fyne.ab/src/core/gui"
	"sync"
)

var (
	eApi  *EmptyPage
	eOnce sync.Once
)

func ApiEmptyPage() *EmptyPage {
	eOnce.Do(func() {
		eApi = new(EmptyPage)
	})
	return eApi
}

type EmptyPage struct {
	BasePage
}

func (tp *EmptyPage) Init(window *gui.AppWindow) gui.IPage {
	tp.window = window
	tp.id = snowman.NewSnowApi().GetIdInt64()
	tp.name = "空白页面"
	tp.flagIsError = false
	tp.flagIsStart = false

	return tp
}

func (tp *EmptyPage) GenCanvasObject() fyne.CanvasObject {
	return container.New(layout.NewMaxLayout())
}
