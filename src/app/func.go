package app

import (
	"fmt"
	"fyne.io/fyne/v2"
	"gui.fyne.ab/src/app/page"
	"gui.fyne.ab/src/core/gui"
)

func genWindow() *gui.AppWindow {
	MainWindow.SetMainMenu(genMenu())

	MainWindow.RegisterPages(page.ApiTestPage().Init(MainWindow))
	MainWindow.RegisterPages(page.ApiEmptyPage().Init(MainWindow))
	MainWindow.RegisterPages(page.ApiErrorPage().Init(MainWindow))

	return MainWindow
}

func genMenu() *fyne.MainMenu {
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu(
			"菜单",
			fyne.NewMenuItem("设置", func() {}),
		),
		fyne.NewMenu(
			"服务",
			fyne.NewMenuItem("AB压测", func() {

			}),
		),
		fyne.NewMenu(
			"帮助",
			fyne.NewMenuItem("检测版本", func() {
				fmt.Println("点击检测版本")
			}),
			fyne.NewMenuItem("关于应用", func() {
				fmt.Println("点击关于应用")
			}),
		),
		fyne.NewMenu(
			"页面",
			fyne.NewMenuItem("测试页面", func() {
				MainWindow.SetCurrentPageId(page.ApiTestPage().GetId())
			}),
			fyne.NewMenuItem("空白页面", func() {
				MainWindow.SetCurrentPageId(page.ApiEmptyPage().GetId())
			}),
			fyne.NewMenuItem("错误页面", func() {
				MainWindow.SetCurrentPageId(page.ApiErrorPage().GetId())
			}),
		),
	)
	return mainMenu
}
