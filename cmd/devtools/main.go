package main

import (
	"devtools/internal/app"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Создаем и запускаем приложение
	app := app.NewApp()
	p := tea.NewProgram(app, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
