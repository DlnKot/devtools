package models

import (
	"devtools/internal/tools"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type PingModel struct {
	BaseModel
	input   textinput.Model
	output  string
	loading bool
	target  string
}

func NewPingModel() *PingModel {
	ti := textinput.New()
	ti.Placeholder = "Enter hostname or IP address (e.g., google.com or 8.8.8.8)"
	ti.Focus()
	ti.CharLimit = 255
	ti.Width = 50

	return &PingModel{
		input: ti,
	}
}

func (m *PingModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *PingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "esc":
			// Возвращаемся к NetworkModel
			return m, func() tea.Msg {
				return SwitchModelMsg{Model: NewNetworkModel()}
			}

		case "enter":
			if m.input.Focused() {
				m.target = m.input.Value()
				if m.target != "" {
					m.loading = true
					m.output = ""
					// Запускаем ping
					return m, m.pingCommand(m.target)
				}
			}
		}

	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)

	case PingResultMsg:
		m.loading = false
		m.output = string(msg)
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

type PingResultMsg string

// Команда для выполнения ping
func (m *PingModel) pingCommand(target string) tea.Cmd {
	return func() tea.Msg {
		pingTools := tools.PingTools{}
		result, err := pingTools.Ping()
		if err != nil {
			return PingResultMsg("Error ping")
		}
		return PingResultMsg(result)
	}
}

func (m *PingModel) View() string {
	var b strings.Builder

	b.WriteString("🏓 Ping Utility\n\n")
	b.WriteString("Target: ")
	b.WriteString(m.input.View())
	b.WriteString("\n\n")

	if m.loading {
		b.WriteString("⏳ Pinging ")
		b.WriteString(m.target)
		b.WriteString("...\n")
	} else if m.output != "" {
		// b.WriteString("📊 Results:\n")
		b.WriteString(m.output)
		b.WriteString("\n\n")
	}

	b.WriteString("\n\nEnter: Start ping • esc: Back to Network • q: Quit")

	return b.String()
}
