package modules

import (
	"fyne.io/fyne/v2"
)

type Module interface {
	ID() string
	Name() string

	OnEnable()
	OnDisable()

	GetView() fyne.CanvasObject

	IsEnable() bool
}

type ModulesManager struct {
	modules []Module
	enabled []Module
}

func NewModulesManager(modules ...Module) ModulesManager {
	return ModulesManager{
		modules: modules,
	}
}

func (m *ModulesManager) RegisterModule(module Module) {
	m.modules = append(m.modules, module)
	if module.IsEnable() {
		module.OnEnable() // Автоматически включаем модуль при регистрации
	}
}

func (m ModulesManager) GetModules() []Module {
	return m.enabled
}
