package menu

import (
	"fmt"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type MenuItemPressedMsg struct {
	ID string
}

type MenuItem struct {
	Title string
	ID    string
}

type Menu struct {
	Items      []MenuItem
	zManager   *zone.Manager
	hoverIndex int
}

func NewMenu(items []MenuItem) Menu {
	return Menu{
		Items:      items,
		zManager:   zone.New(),
		hoverIndex: 0,
	}
}

func (m Menu) Init() tea.Cmd {
	return nil
}

func (m Menu) Update(msg tea.Msg) (Menu, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if msg.Action == tea.MouseActionMotion {
			for index := range m.Items {
				if m.zManager.Get(fmt.Sprint(index)).InBounds(msg) {
					m.hoverIndex = index
					return m, nil
				}
			}
			break
		}

		if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft {
			return m, nil
		}

		for index, i := range m.Items {
			if m.zManager.Get(fmt.Sprint(index)).InBounds(msg) {
				return m, func() tea.Msg {
					return MenuItemPressedMsg{ID: i.ID}
				}
			}
		}
	}

	return m, nil
}

func (m Menu) View() string {
	longestWidth := 0
	for _, s := range m.Items {
		lenght := utf8.RuneCount([]byte(s.Title))
		if lenght > longestWidth {
			longestWidth = lenght
		}
	}

	items := []string{}
	for index, i := range m.Items {
		var item string
		if m.hoverIndex == index {
			item = MenuItemActiveStyle.Width(longestWidth).Render(i.Title)
		} else {
			item = MenuItemStyle.Width(longestWidth).Render(i.Title)
		}
		items = append(items, m.zManager.Mark(fmt.Sprint(index), item))
	}

	menu := lipgloss.JoinVertical(lipgloss.Top, items...)
	return m.zManager.Scan(MenuStyle.Render(menu))
}
