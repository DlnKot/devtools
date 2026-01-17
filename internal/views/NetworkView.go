package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type NetworkView struct {
}

func (v NetworkView) init() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("Ping", widget.NewLabel("Pingsdsd")),
		container.NewTabItem("API req", widget.NewLabel("dfd")),
	)
	tabs.SetTabLocation(container.TabLocationBottom)

	return tabs
}
