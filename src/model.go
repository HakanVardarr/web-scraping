package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	categories []category
	category   *category
	cursor     int
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
			m.category = &categories()[m.cursor]
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.categories)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.category == nil {
		content := "Welcome to Trendyol!\n\n"
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
		content := "Category: "
		content += m.category.Name
		content += "\n\n"

		content += "\nPress q to quit.\n"

		return content
	}
}
