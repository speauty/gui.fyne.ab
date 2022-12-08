package page

import (
	"gui.fyne.ab/src/core/gui"
)

type BasePage struct {
	Id          int64
	Name        string
	FlagIsError bool
	FlagIsStart bool
	Window      *gui.AppWindow
}

func (basePage *BasePage) GetId() int64 { return basePage.Id }

func (basePage *BasePage) GetName() string { return basePage.Name }

func (basePage *BasePage) FlagIsErrorPage() bool { return basePage.FlagIsError }

func (basePage *BasePage) FlagIsStartPage() bool { return basePage.FlagIsStart }
