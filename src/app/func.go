package app

import (
	"fyne.io/fyne/v2"
	"gui.fyne.ab/src/app/page"
	"gui.fyne.ab/src/core/gui"
)

func genWindow() *gui.AppWindow {
	MainWindow.SetMainMenu(genMenu())

	MainWindow.RegisterPages(page.ApiErrorPage().Init(MainWindow))
	MainWindow.RegisterPages(page.ApiEnvPage().Init(MainWindow))
	MainWindow.RegisterPages(page.ApiABPage().Init(MainWindow))

	return MainWindow
}

func genMenu() *fyne.MainMenu {
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu(
			"菜单",
			fyne.NewMenuItem("环境", func() {
				go LogAppApi().Infoln("点击菜单: 菜单-环境")
			}),
			fyne.NewMenuItem("设置", func() {
				go LogAppApi().Infoln("点击菜单: 菜单-设置")
			}),
		),
		fyne.NewMenu(
			"服务",
			fyne.NewMenuItem("AB压测", func() {
				go LogAppApi().Infoln("点击菜单: 服务-AB压测")
				MainWindow.LoadPage(page.ApiABPage().GetId())
			}),
		),
		fyne.NewMenu(
			"页面",
			fyne.NewMenuItem("错误页面", func() {
				go LogAppApi().Infoln("点击菜单: 页面-错误页面")
				MainWindow.LoadPage(page.ApiErrorPage().GetId())
			}),
		),
	)
	return mainMenu
}
