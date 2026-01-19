package app

import (
	"devtools/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	currentModel tea.Model
}

func NewApp() *App {
	
	return &App{
		currentModel: models.NewMenuModel(),
	}
}

func (a *App) Init() tea.Cmd {
	return a.currentModel.Init()
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Обновляем текущую модель
	a.currentModel, cmd = a.currentModel.Update(msg)

	// Проверяем, не нужно ли переключить модель
	if switchMsg, ok := msg.(models.SwitchModelMsg); ok {
		a.currentModel = switchMsg.Model
		cmd = a.currentModel.Init()
	}

	return a, cmd
}

func (a *App) View() string {
	return a.currentModel.View()
}
