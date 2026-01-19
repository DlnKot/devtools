package models

import tea "github.com/charmbracelet/bubbletea"

// Сообщение для переключения между моделями
type SwitchModelMsg struct {
	Model tea.Model
}

// Базовая структура для всех моделей
type BaseModel struct {
	Width  int
	Height int
}

func (b *BaseModel) SetSize(width, height int) {
	b.Width = width
	b.Height = height
}
