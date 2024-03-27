package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	width, height            int
	categories               []category
	category                 *category
	cursor, cursorSave, page int
}

func newModel(categories []category) model {
	return model{0, 0, categories, nil, 0, 0, 1}

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
			if m.category == nil {
				m.category = &m.categories[m.cursor]
				m.cursorSave = m.cursor
				m.cursor = 0

			}
		case "esc":
			if m.category != nil {
				m.category = nil
				m.cursor = m.cursorSave
			}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right":
			if m.category != nil {
				m.page += 1
				m.category.Products.Clear()
			}
		case "left":
			if m.category != nil {
				if m.page > 1 {
					m.page -= 1
					m.category.Products.Clear()
				}
			}

		case "down":
			if m.category != nil {
				if m.cursor < m.category.Products.Length()-1 {
					m.cursor++
				}
			} else {
				if m.cursor < len(m.categories)-1 {
					m.cursor++
				}
			}

		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	return m, nil
}

func (m model) View() string {
	return m.category.View(m)
}
