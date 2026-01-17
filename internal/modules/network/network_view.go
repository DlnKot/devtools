package network

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"devtools/internal/modules"
)

type NetworkView struct {
	controller modules.Module
}

func NewNetworkView(controller modules.Module) NetworkView {
	return NetworkView{
		controller: controller,
	}
}

func (v NetworkView) CreateView() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItem("Ping", widget.NewLabel("Hello")),
	)

	return tabs
}
