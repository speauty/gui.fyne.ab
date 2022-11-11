package page

import (
	"gui.fyne.ab/src/core/gui"
)

type BasePage struct {
	id          int64
	name        string
	flagIsError bool
	flagIsStart bool
	window      *gui.AppWindow
}

func (basePage *BasePage) GetId() int64 { return basePage.id }

func (basePage *BasePage) GetName() string { return basePage.name }

func (basePage *BasePage) FlagIsErrorPage() bool { return basePage.flagIsError }

func (basePage *BasePage) FlagIsStartPage() bool { return basePage.flagIsStart }
