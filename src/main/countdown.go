package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	// 创建应用
	myApp := app.New()
	myWindow := myApp.NewWindow("倒计时")

	// 倒计时时间（例如：5分钟）
	countdownTime := 5 * time.Minute
	// 倒计时文本
	label := widget.NewLabelWithStyle("剩余时间: 5:00", fyne.TextAlignCenter, fyne.TextStyle{})
	// 开始按钮
	startButton := widget.NewButton("开始倒计时", func() {
		// 启动倒计时
		updateLabel(myWindow, label, countdownTime)
	})

	// 将组件放入容器
	myWindow.SetContent(container.NewVBox(label, startButton))

	// 显示并运行应用
	myWindow.ShowAndRun()
}

func updateLabel(myWindow fyne.Window, label *widget.Label, countdownTime time.Duration) {
	myWindow.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		if ke.Name == fyne.KeyEscape {
			countdownTime = 0
		}
	})
	if countdownTime > 0 {
		label.SetText("剩余时间: " + countdownTime.String())
		countdownTime = countdownTime - time.Second
		time.AfterFunc(time.Second, func() {
			updateLabel(myWindow, label, countdownTime)
		})
	} else {
		label.SetText("倒计时结束!")
		return
	}
}
