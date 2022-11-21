package app

import (
	"fyne.io/fyne/v2"
	"gui.fyne.ab/src/app/page"
	"gui.fyne.ab/src/core/gui"
)

func genWindow() *gui.AppWindow {
	MainWindow.SetMainMenu(genMenu())

	MainWindow.RegisterPages(page.ApiTestPage().Init(MainWindow))
	MainWindow.RegisterPages(page.ApiEmptyPage().Init(MainWindow))
	MainWindow.RegisterPages(page.ApiErrorPage().Init(MainWindow))
	MainWindow.RegisterPages(page.ApiEnvPage().Init(MainWindow))

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
			fyne.NewMenuItem("上门采样", func() {
				go LogAppApi().Infoln("点击菜单: 服务-上门采样")
			}),
		),
		fyne.NewMenu(
			"帮助",
			fyne.NewMenuItem("检测版本", func() {
				go LogAppApi().Infoln("点击菜单: 帮助-检测版本")
			}),
			fyne.NewMenuItem("关于应用", func() {
				go LogAppApi().Infoln("点击菜单: 帮助-关于应用")
			}),
		),
		fyne.NewMenu(
			"页面",
			fyne.NewMenuItem("测试页面", func() {
				go LogAppApi().Infoln("点击菜单: 页面-测试页面")
				MainWindow.SetCurrentPageId(page.ApiTestPage().GetId())
			}),
			fyne.NewMenuItem("空白页面", func() {
				go LogAppApi().Infoln("点击菜单: 页面-空白页面")
				MainWindow.SetCurrentPageId(page.ApiEmptyPage().GetId())
			}),
			fyne.NewMenuItem("错误页面", func() {
				go LogAppApi().Infoln("点击菜单: 页面-错误页面")
				MainWindow.SetCurrentPageId(page.ApiErrorPage().GetId())
			}),
		),
	)
	return mainMenu
}
