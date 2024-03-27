package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	container      = lipgloss.NewStyle().Bold(true).Padding(2)
	defaultText    = lipgloss.NewStyle().Foreground(lipgloss.Color("#d3bacd"))
	chosenCategory = lipgloss.NewStyle().Foreground(lipgloss.Color("#d315a4"))
	price          = lipgloss.NewStyle().Foreground(lipgloss.Color("#04ed18"))
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
				m.category.Items = make([]item, 0)
			}
		case "left":
			if m.category != nil {
				if m.page > 1 {
					m.page -= 1
					m.category.Items = make([]item, 0)
				}
			}

		case "down":
			if m.category != nil {
				if m.cursor < len(m.category.Items)-1 {
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
	if m.category == nil {

		content := fmt.Sprintf("%s\n\n", defaultText.Render("Choose your category"))

		for i, category := range m.categories {
			cursor := " "

			if m.cursor == i {
				cursor = "> "
				content += fmt.Sprintf("%s\n", chosenCategory.Render(cursor+category.Name))
			} else {
				content += fmt.Sprintf("%s\n", defaultText.Render(category.Name))
			}

		}

		content += fmt.Sprintf("\n%s", defaultText.Render("Press Q to quit!"))

		return container.Render(content)
	} else {
		content := fmt.Sprintf("%s\n\n", defaultText.Render("Press ESC to return category selection."))
		content += fmt.Sprintf("%s%s\n", defaultText.Render("Listing: "), chosenCategory.Bold(true).Render(m.category.Name))
		content += fmt.Sprintf("%s%s\n\n", defaultText.Render("Page: "), chosenCategory.Bold(true).Render(fmt.Sprint(m.page)))

		if len(m.category.Items) == 0 {
			resp, err := soup.Get(m.category.Link + fmt.Sprint(m.page))

			if err != nil {
				os.Exit(1)
			}

			doc := soup.HTMLParse(resp)
			products := doc.FindAll("div", "class", "p-card-wrppr")

			for i, product := range products {
				productName := product.Find("span", "class", "prdct-desc-cntnr-name").Text()
				productPrice := product.Find("div", "class", "prc-box-dscntd").Text()

				cursor := " "

				if m.cursor == i {
					cursor = "> "
					content += fmt.Sprintf("%s %s\n", chosenCategory.Render(cursor+productName), price.Render(productPrice))
				} else {
					content += fmt.Sprintf("%s %s\n", defaultText.Render(productName), defaultText.Render(productPrice))
				}

				m.category.Items = append(m.category.Items, item{productName, productPrice})
			}

		} else {
			for i, product := range m.category.Items {
				cursor := " "

				if m.cursor == i {
					cursor = "> "
					content += fmt.Sprintf("%s %s\n", chosenCategory.Render(cursor+product.Name), price.Render(product.Price))
				} else {
					content += fmt.Sprintf("%s %s\n", defaultText.Render(product.Name), defaultText.Render(product.Price))
				}
			}
		}

		content += fmt.Sprintf("\n%s", defaultText.Render("Press Q to quit!"))

		return container.Render(content)
	}
}
