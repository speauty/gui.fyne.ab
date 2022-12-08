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
	errorPageApi  *ErrorPage
	errorPageOnce sync.Once
)

func ApiErrorPage() *ErrorPage {
	errorPageOnce.Do(func() {
		errorPageApi = new(ErrorPage)
	})
	return errorPageApi
}

type ErrorPage struct {
	BasePage
}

func (errorPage *ErrorPage) Init(window *gui.AppWindow) gui.IPage {
	errorPage.Window = window
	errorPage.Id = snowman.NewSnowApi().GetIdInt64()
	errorPage.Name = "错误页面"
	errorPage.FlagIsError = true
	return errorPage
}

func (errorPage *ErrorPage) GenCanvasObject() fyne.CanvasObject {
	return container.New(layout.NewCenterLayout(), widget.NewLabel("当前应用发生异常"))
}
