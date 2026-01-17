package app

import (
	"devtools/internal/modules"
	"devtools/internal/modules/network"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type App struct {
	fyneApp        fyne.App
	Window         fyne.Window
	ModulesManager modules.ModulesManager
	Views          []fyne.Container
}

func New() *App {
	a := app.New()
	w := a.NewWindow("Dvtls")
	w.Resize(fyne.NewSize(800, 600))

	app := &App{
		fyneApp: a,
		Window:  w,
	}

	// TODO: Реализовать выбор модулей через конфиг
	app.ModulesManager = modules.NewModulesManager()
	app.ModulesManager.RegisterModule(network.NewNetworkModule("network", "Network", true))

	return app
}

func (a *App) Run() {
	// Создаем вкладки для каждого модуля
	tabs := container.NewAppTabs()

	for _, module := range a.ModulesManager.GetModules() {
		tabs.Append(container.NewTabItem(module.Name(), module.GetView()))
	}

	a.Window.SetContent(tabs)
	a.Window.ShowAndRun()
}
