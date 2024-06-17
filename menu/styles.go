package menu

import (
	"github.com/mluna-again/the_tea/internal"

	"github.com/charmbracelet/lipgloss"
)

var MenuStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderLeft(false).
	BorderTop(false).
	BorderForeground(internal.Black).
	BorderBackground(internal.Black)

var MenuItemStyle = lipgloss.NewStyle().
	Background(internal.Blue).
	Foreground(internal.White)

var MenuItemActiveStyle = lipgloss.NewStyle().
	Background(internal.White).
	Foreground(internal.Black)
