package tools

import tea "github.com/charmbracelet/bubbletea"

type ToolsItem struct {
	Name        string
	Description string
	Model       func() tea.Model
}
