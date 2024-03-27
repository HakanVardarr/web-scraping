package main

import "github.com/charmbracelet/lipgloss"

var (
	container      = lipgloss.NewStyle().Padding(2)
	defaultText    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#d3bacd"))
	chosenCategory = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#d315a4"))
	price          = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#04ed18"))
)
