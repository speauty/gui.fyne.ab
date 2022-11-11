package gui

import "fyne.io/fyne/v2"

type IPage interface {
	Init(window *AppWindow) IPage
	GetId() int64
	GetName() string
	FlagIsErrorPage() bool // 是否错误页面
	FlagIsStartPage() bool // 是否启动页面
	GenCanvasObject() fyne.CanvasObject
}
