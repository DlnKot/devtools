package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainView struct {
	Width  float32
	Height float32
}

func (w MainView) Initialize(window fyne.Window) fyne.Window {

	networkView := NetworkView{}.init()

	tabs := container.NewAppTabs(
		container.NewTabItem("Network", networkView),
		container.NewTabItem("Passwords", widget.NewLabel("Hello")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	window.Resize(fyne.NewSize(w.Width, w.Height))

	window.SetContent(
		tabs,
	)

	return window
}
