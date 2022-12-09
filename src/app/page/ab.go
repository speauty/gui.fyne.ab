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
	ap.Window = window
	ap.Id = snowman.NewSnowApi().GetIdInt64()
	ap.Name = "AB压测"
	ap.FlagIsError = false
	ap.FlagIsStart = true
	ap.exePath = filepath.Join(gui.GetRuntimeDir(), "ab.exe")

	return ap
}

func (ap *ABPage) GenCanvasObject() fyne.CanvasObject {
	itemTargetLink := new(widget.FormItem)
	itemTargetLinkEntry := widget.NewEntry()
	itemTargetLinkEntry.SetText("http://www.baidu.com/")
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
	itemCntRequestEntry.SetText("1")
	itemCntRequestEntry.Resize(fyne.NewSize(80, 40))
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

	itemCntParallel := new(widget.FormItem)
	itemCntParallelEntry := widget.NewEntry()
	itemCntParallelEntry.SetText("1")
	itemCntParallelEntry.Validator = func(val string) error {
		if val == "" {
			return fmt.Errorf("请输入并发数量")
		}
		numParsed, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("并发数量解析异常, 错误: %s", err.Error())
		}
		if strconv.Itoa(numParsed) != val {
			return fmt.Errorf("并发数量解析不一致, 输入值: %s, 解析值: %s", val, strconv.Itoa(numParsed))
		}
		return nil
	}
	itemCntParallel.Text = "并发数量"
	itemCntParallel.Widget = itemCntParallelEntry

	itemTimeExecuted := new(widget.FormItem)
	itemTimeExecutedEntry := widget.NewEntry()
	itemTimeExecutedEntry.SetText("10")
	itemTimeExecutedEntry.Validator = func(val string) error {
		if val == "" {
			return fmt.Errorf("请输入运行时间")
		}
		numParsed, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("运行时间解析异常, 错误: %s", err.Error())
		}
		if strconv.Itoa(numParsed) != val {
			return fmt.Errorf("运行时间解析不一致, 输入值: %s, 解析值: %s", val, strconv.Itoa(numParsed))
		}
		return nil
	}
	itemTimeExecuted.Text = "运行时间(s)"
	itemTimeExecuted.Widget = itemTimeExecutedEntry

	strRes := ""
	dataRes := binding.BindString(&strRes)
	dataScreen := widget.NewLabelWithData(dataRes)

	mainForm := new(widget.Form)
	mainForm.SubmitText = "发送"
	mainForm.CancelText = "重置"
	mainForm.OnSubmit = func() {
		_ = dataRes.Set("")
		var args []string
		if itemCntRequestEntry.Text != "" {
			args = append(args, "-n", itemCntRequestEntry.Text)
		}
		if itemCntParallelEntry.Text != "" {
			args = append(args, "-c", itemCntParallelEntry.Text)
		}
		if itemTimeExecutedEntry.Text != "" {
			args = append(args, "-t", itemTimeExecutedEntry.Text)
		}
		args = append(args, itemTargetLinkEntry.Text)
		cmd := exec.Command(ap.exePath, args...)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println(cmd.Args)
			gui.Error(err, ap.Window.GetWindow())
		}
		_ = dataRes.Set(out.String())
		return
	}
	mainForm.OnCancel = func() {
		itemTargetLinkEntry.SetText("http://www.baidu.com/")
		itemCntRequestEntry.SetText("1")
		itemCntParallelEntry.SetText("1")
		itemTimeExecutedEntry.SetText("10")
		_ = dataRes.Set("")
	}

	mainForm.Items = append(mainForm.Items, itemTargetLink, itemCntRequest, itemCntParallel, itemTimeExecuted)
	windowSize := ap.Window.GetWindow().Content().Size()
	if windowSize.Width == 0 {
		windowSize = fyne.NewSize(800, 600)
	}
	return container.NewMax(
		container.NewVBox(
			container.New(layout.NewGridWrapLayout(fyne.NewSize(windowSize.Width, windowSize.Height-35)), container.NewMax(container.NewHBox(
				container.New(layout.NewGridWrapLayout(fyne.NewSize(windowSize.Width*0.65, windowSize.Height-35)), container.NewMax(container.NewHScroll(mainForm))),
				container.New(layout.NewGridWrapLayout(fyne.NewSize(windowSize.Width-windowSize.Width*0.65, windowSize.Height-35)), container.NewMax(container.NewScroll(dataScreen))),
			))),
			container.New(layout.NewGridWrapLayout(fyne.NewSize(windowSize.Width, 35)), container.NewMax(container.NewHScroll(widget.NewLabel("状态栏：fsaf\\n粉红色卡夫卡撒、、\\n核辐射恐慌饭卡")))),
		),
	)
}
