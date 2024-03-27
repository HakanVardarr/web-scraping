package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cats, err := categories()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := newModel(cats)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}
