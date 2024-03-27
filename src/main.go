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

	cats, err := categories()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := model{
		0, 0,
		cats,
		nil,
		0,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}
