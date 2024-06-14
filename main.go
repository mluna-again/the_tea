package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const exitBtnID = "exit"

type model struct {
	menu          Menu
	buttonPressed string
	buttons       []MenuItem
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case MenuItemPressedMsg:
		m.buttonPressed = msg.ID
		if m.buttonPressed == exitBtnID {
			return m, tea.Quit
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.menu, cmd = m.menu.Update(msg)

	return m, cmd
}

func (m model) View() string {
	menu := m.menu.View()

	button := ""
	for _, b := range m.buttons {
		if b.ID == m.buttonPressed {
			button = b.Title
		}
	}

	if button == "" {
		return menu
	}

	return lipgloss.JoinVertical(lipgloss.Top, menu, fmt.Sprintf("You pressed the %s button", button))
}

func newModel(menu Menu) model {
	return model{
		menu:          menu,
		buttons:       menu.Items,
		buttonPressed: "",
	}
}

func main() {
	items := []MenuItem{
		{
			Title: "Do Nothing Button",
			ID:    "1",
		},
		{
			Title: "About this program",
			ID:    "2",
		},
		{
			Title: "Help",
			ID:    "3",
		},
		{
			Title: "Preferences",
			ID:    "4",
		},
		{
			Title: "Exit",
			ID:    exitBtnID,
		},
	}
	menu := NewMenu(items)

	p := tea.NewProgram(newModel(menu), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
