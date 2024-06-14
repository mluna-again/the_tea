package menu

import (
	"fmt"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	zone "github.com/lrstanley/bubblezone"
)

type MenuItemPressedMsg struct {
	ID string
}

type MenuItem struct {
	Title   string
	ID      string
	Submenu *Menu
}

type Menu struct {
	Items      []MenuItem
	Active     bool
	Root       bool
	zManager   *zone.Manager
	hoverIndex int
	id         string
}

func NewMenu(items []MenuItem, root bool, zone *zone.Manager) *Menu {
	for i := range items {
		if items[i].Submenu != nil {
			items[i].Title = fmt.Sprintf("%s >", items[i].Title)
		}
	}

	return &Menu{
		Items:      items,
		zManager:   zone,
		hoverIndex: 0,
		Active:     root,
		Root:       root,
		id:         uuid.NewString(),
	}
}

func (m Menu) Init() tea.Cmd {
	return nil
}

func (m Menu) Update(msg tea.Msg) (Menu, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if !m.zManager.Get(m.id).InBounds(msg) {
			break
		}

		if msg.Action == tea.MouseActionMotion {
			for index, i := range m.Items {
				if m.zManager.Get(i.ID).InBounds(msg) {
					m.hoverIndex = index
					break
				}
			}
			break
		}

		if msg.Action != tea.MouseActionRelease || msg.Button != tea.MouseButtonLeft || !m.Active {
			break
		}

		for _, i := range m.Items {
			clicked := m.zManager.Get(i.ID).InBounds(msg)
			if clicked && i.Submenu == nil {
				m.disableOtherMenus("")
				return m, func() tea.Msg {
					return MenuItemPressedMsg{ID: i.ID}
				}
			}

			if clicked && i.Submenu != nil {
				i.Submenu.Active = true
				m.disableOtherMenus(i.ID)
				break
			}

			if clicked {
				// disable all menus
				m.disableOtherMenus("")
			}
		}
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd
	for i, item := range m.Items {
		if item.Submenu != nil && item.Submenu.Active {
			var newMenu Menu
			newMenu, cmd = m.Items[i].Submenu.Update(msg)
			m.Items[i].Submenu = &newMenu
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Menu) View() string {
	longestWidth := 0
	for _, s := range m.Items {
		lenght := utf8.RuneCount([]byte(s.Title))
		if lenght > longestWidth {
			longestWidth = lenght
		}
	}

	submenus := []string{}
	items := []string{}

	for index, i := range m.Items {
		var item string
		if m.hoverIndex == index {
			item = MenuItemActiveStyle.Width(longestWidth).Render(i.Title)
		} else {
			item = MenuItemStyle.Width(longestWidth).Render(i.Title)
		}
		items = append(items, m.zManager.Mark(fmt.Sprint(i.ID), item))

		if i.Submenu != nil && i.Submenu.Active {
			submenus = append(submenus, i.Submenu.View())
		}
	}

	menu := lipgloss.JoinVertical(lipgloss.Top, items...)
	menu = MenuStyle.Render(menu)
	allMenus := []string{menu}
	allMenus = append(allMenus, submenus...)

	content := lipgloss.JoinHorizontal(lipgloss.Center, allMenus...)

	return m.zManager.Mark(m.id, content)
}

func (m *Menu) disable() {
	m.Active = false
}

func (m *Menu) disableOtherMenus(id string) {
	for _, item := range m.Items {
		if item.Submenu != nil && (item.ID != id || id == "") {
			item.Submenu.disable()
			item.Submenu.disableOtherMenus(id)
		}
	}
}
