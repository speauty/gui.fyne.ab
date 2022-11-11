package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"gui.fyne.ab/src/core/gui"
	"image/color"
	"math/rand"
	"sync"
)

var (
	tApi  *TestPage
	tOnce sync.Once
)

func ApiTestPage() *TestPage {
	tOnce.Do(func() {
		tApi = new(TestPage)
	})
	return tApi
}

type TestPage struct {
	BasePage
}

func (tp *TestPage) Init(window *gui.AppWindow) gui.IPage {
	tp.window = window
	tp.id = rand.Int63()
	tp.name = "测试页面"
	tp.flagIsError = false
	tp.flagIsStart = true

	return tp
}

func (tp *TestPage) GenCanvasObject() fyne.CanvasObject {
	img := canvas.NewImageFromResource(theme.FyneLogo())
	text := canvas.NewText(gui.Api().GetAppName(), color.White)
	content := container.New(layout.NewMaxLayout(), img, container.NewCenter(text))

	return content
}
