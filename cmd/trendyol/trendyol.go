package main

import (
	"fmt"
	"os"

	"github.com/HakanVardarrhakan/web-scraping/pkg/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	categories, err := model.Categories()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := model.New(categories)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}
