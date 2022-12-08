package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func Error(err error, window fyne.Window) dialog.Dialog {
	alert := dialog.NewError(err, window)
	alert.SetDismissText("确认")
	alert.Show()
	return alert
}
