package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width, height int
	categories    []category
	category      *category
	cursor        int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.category = &m.categories[m.cursor]
		case "esc":
			if m.category != nil {
				m.category = nil
			}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.categories)-1 {
				m.cursor++
			}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	return m, nil
}

func (m model) View() string {
	if m.category == nil {
		header := lipgloss.NewStyle().Background(lipgloss.Color("#fff")).Width(m.width).Padding(1).Render
		headerText := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#fff")).Foreground(lipgloss.Color("#000")).Inline(true).Render

		content := fmt.Sprintln(header(fmt.Sprintf("%s%s", headerText("Welcome to "), headerText("Trendyol"))))

		content += "Choose your category\n\n"

		for i, category := range m.categories {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			content += fmt.Sprintf("%s %s\n", cursor, category.Name)
		}

		content += "\nPress q to quit.\n"

		return content
	} else {
		content := "Press ESC to return category selection.\n\nCategory: "
		content += m.category.Name
		content += "\n\n"

		content += "\nPress q to quit.\n"

		return content
	}
}
