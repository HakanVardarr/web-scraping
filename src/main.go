package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	URL = "https://www.trendyol.com/"
)

func main() {

	m := model{
		categories(),
		nil,
		0,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}
