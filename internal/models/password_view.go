package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type PassModel struct {
	BaseModel
	cursor  int
	choices []MenuItem
}

func NewPassModel() *PassModel {
	return &PassModel{
		choices: []MenuItem{
			{
				Name:  "Generate password",
				// Model: func() tea.Model { return NewPingModel() },
			},
			{
				Name: "Check pass",
				// Model:       tools.NewBase64Encoder,
			},
		},
	}
}

func (m *PassModel) Init() tea.Cmd {
	return nil
}

func (m *PassModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter":
			// Создаем и переключаемся на выбранную утилиту
			selectedTool := m.choices[m.cursor].Model()
			return m, func() tea.Msg {
				return SwitchModelMsg{Model: selectedTool}
			}

		case "esc":
			// Возвращаемся в главное меню
			return m, func() tea.Msg {
				return SwitchModelMsg{Model: NewMenuModel()}
			}
		}

	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
	}

	return m, nil
}

func (m *PassModel) View() string {
	s := "🚀 Network - Select a utility\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = "▶"
		}

		s += fmt.Sprintf("%s  %s\n",
			cursor,
			choice.Name,
		)
	}

	s += "\n\n↑/↓: Navigate • Enter: Select • q: Quit • esc: Back"
	return s
}
