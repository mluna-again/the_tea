package menu

import "github.com/charmbracelet/lipgloss"

var black = lipgloss.Color("#181816")
var white = lipgloss.Color("#c5c9c5")
var gray = lipgloss.Color("#a6a69c")

var MenuStyle = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder()).
	BorderForeground(white)

var MenuItemStyle = lipgloss.NewStyle().
	Background(black).
	Foreground(white)

var MenuItemInactiveStyle = lipgloss.NewStyle().
	Background(gray).
	Foreground(white)
