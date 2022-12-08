package page

import (
	"bytes"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gui.fyne.ab/src/common/snowman"
	"gui.fyne.ab/src/core/gui"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
)

var (
	abApi  *ABPage
	abOnce sync.Once
)

func ApiABPage() *ABPage {
	abOnce.Do(func() {
		abApi = new(ABPage)
	})
	return abApi
}

type ABPage struct {
	BasePage
	exePath string
}

func (ap *ABPage) Init(window *gui.AppWindow) gui.IPage {
	ap.window = window
	ap.id = snowman.NewSnowApi().GetIdInt64()
	ap.name = "AB压测"
	ap.flagIsError = false
	ap.flagIsStart = false
	ap.exePath = filepath.Join(gui.GetRuntimeDir(), "ab.exe")

	return ap
}

func (ap *ABPage) GenCanvasObject() fyne.CanvasObject {

	itemTargetLink := new(widget.FormItem)
	itemTargetLinkEntry := widget.NewEntry()
	itemTargetLinkEntry.Validator = func(val string) error {
		if val == "" {
			return fmt.Errorf("请输入目标链接地址")
		}
		// @todo 可插入正则验证, 无效链接
		return nil
	}
	itemTargetLink.Text = "目标链接"
	itemTargetLink.Widget = itemTargetLinkEntry

	itemCntRequest := new(widget.FormItem)
	itemCntRequestEntry := widget.NewEntry()
	itemCntRequestEntry.SetText("10")
	itemCntRequestEntry.Validator = func(val string) error {
		if val == "" {
			return fmt.Errorf("请输入总请求数")
		}
		numParsed, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("总请求数解析异常, 错误: %s", err.Error())
		}
		if strconv.Itoa(numParsed) != val {
			return fmt.Errorf("总请求数解析不一致, 输入值: %s, 解析值: %s", val, strconv.Itoa(numParsed))
		}
		return nil
	}
	itemCntRequest.Text = "总请求数"
	itemCntRequest.Widget = itemCntRequestEntry

	strRes := ""
	dataRes := binding.BindString(&strRes)
	dataScreen := widget.NewLabelWithData(dataRes)

	mainForm := new(widget.Form)
	mainForm.SubmitText = "发送"
	mainForm.CancelText = "重置"
	mainForm.OnSubmit = func() {
		_ = dataRes.Set("")
		var args []string
		args = append(args, "-n", itemCntRequestEntry.Text)
		args = append(args, itemTargetLinkEntry.Text)
		cmd := exec.Command(ap.exePath, args...)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			gui.Error(err, ap.window.GetWindow())
		}
		_ = dataRes.Set(out.String())
		return
	}
	mainForm.OnCancel = func() {
		itemTargetLinkEntry.SetText("")
		itemCntRequestEntry.SetText("10")
		_ = dataRes.Set("")
	}
	mainForm.Items = append(mainForm.Items, itemTargetLink, itemCntRequest)
	windowSize := ap.window.GetWindow().Content().Size()
	return container.NewHBox(
		container.New(layout.NewGridWrapLayout(fyne.NewSize(windowSize.Width*0.6, windowSize.Height)), container.NewMax(container.NewHScroll(mainForm))),
		container.New(layout.NewGridWrapLayout(fyne.NewSize(windowSize.Width-windowSize.Width*0.3, windowSize.Height)), container.NewMax(container.NewScroll(dataScreen))),
	)
}
