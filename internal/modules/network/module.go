package network

import "fyne.io/fyne/v2"

type NetworkModule struct {
	id        string
	name      string
	ui        fyne.CanvasObject
	isEnabled bool
	view      NetworkView
}

func NewNetworkModule(id string, name string, isEnabled bool) NetworkModule {
	return NetworkModule{
		id:        id,
		name:      name,
		isEnabled: isEnabled,
	}
}

func (m NetworkModule) GetView() fyne.CanvasObject {
	netWorkView := NewNetworkView(m)
	view := netWorkView.CreateView()
	return view
}

func (m NetworkModule) ID() string {
	return m.id
}

func (m NetworkModule) Name() string {
	return m.name
}

func (m NetworkModule) OnEnable() {
	m.isEnabled = true
}

func (m NetworkModule) OnDisable() {
	m.isEnabled = false
}

func (m NetworkModule) IsEnable() bool {
	return m.isEnabled
}
