package menu

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

const exitBtnID = "exit"

type model struct {
	menu          Menu
	buttonPressed string
	buttons       []MenuItem
	z             *zone.Manager
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
		return m.z.Scan(menu)
	}

	content := lipgloss.JoinVertical(lipgloss.Top, menu, fmt.Sprintf("You pressed the %s button", button))
	return m.z.Scan(content)
}

func newModel(menu *Menu, z *zone.Manager) model {
	return model{
		menu:          *menu,
		buttons:       menu.Items,
		buttonPressed: "",
		z:             z,
	}
}

func Demo() {
	z := zone.New()

	items := []MenuItem{
		{
			Title: "Do Nothing Button",
			ID:    "1",
		},
		{
			Title: "About this program",
			ID:    "2",
			Submenu: NewMenu([]MenuItem{
				{
					Title: "It",
					ID:    "a",
				},
				{
					Title: "Supports",
					ID:    "b",
				},
				{
					Title: "Nested",
					ID:    "c",
				},
				{
					Title: "Menus",
					ID:    "d",
					Submenu: NewMenu([]MenuItem{
						{
							Title: ":D",
							ID:    ":D",
						},
						{
							Title: "D:",
							ID:    "D:",
						},
						{
							Title: ";)",
							ID:    ";)",
						},
					}, false, z),
				},
			}, false, z),
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
			Title: "Another nested menu",
			ID:    "5",
			Submenu: NewMenu([]MenuItem{
				{
					Title: "Hello",
					ID:    "x",
				},
				{
					Title: "new",
					ID:    "y",
				},
				{
					Title: "world",
					ID:    "z",
				},
			}, false, z),
		},
		{
			Title: "Exit",
			ID:    exitBtnID,
		},
	}

	menu := NewMenu(items, true, z)

	p := tea.NewProgram(newModel(menu, z), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
