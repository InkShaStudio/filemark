package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

type IListItem interface {
	GetLabel() string
	GetId() string
}

type List struct {
	Choices  []IListItem
	Cursor   int
	Selected map[int]string
	Multiple bool
}

type Wrapper[T any] struct {
	Data     *T
	getID    func(item *T) string
	getLabel func(item *T) string
}

func (w Wrapper[T]) GetId() string {
	return w.getID(w.Data)
}

func (w Wrapper[T]) GetLabel() string {
	return w.getLabel(w.Data)
}

func WrapItemSlice[T any](list *[]T) []*T {
	result := make([]*T, 0)
	l := *list
	for i, _ := range l {
		result = append(result, &l[i])
	}

	return result
}

func WrapItem[T any](list []*T, getID func(item *T) string, getLabel func(item *T) string) []IListItem {
	result := make([]IListItem, 0)

	for _, item := range list {
		result = append(result, &Wrapper[T]{
			Data:     item,
			getID:    getID,
			getLabel: getLabel,
		})
	}

	return result
}

func InitialListUI(list []IListItem, multiple bool) List {
	return List{
		Choices:  list,
		Cursor:   0,
		Selected: make(map[int]string),
		Multiple: multiple,
	}
}

func (m List) Run() {
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("raw mode error: ", err)
		os.Exit(1)
	}
}

func (m List) Init() tea.Cmd {
	return nil
}

func (m List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	choices := m.Choices
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			} else {
				m.Cursor = len(choices) - 1
			}

		case "down", "j":
			if m.Cursor < len(choices)-1 {
				m.Cursor++
			} else {
				m.Cursor = 0
			}

		case " ":
			_, ok := m.Selected[m.Cursor]

			if m.Multiple {
				if ok {
					delete(m.Selected, m.Cursor)
				} else {
					m.Selected[m.Cursor] = choices[m.Cursor].GetId()
				}
			} else {
				flag := true

				for index := 0; index < len(m.Selected); index++ {
					if _, ok := m.Selected[index]; ok {
						if index == m.Cursor {
							flag = false
						}
						delete(m.Selected, m.Cursor)
					}
				}

				if flag {
					m.Selected[m.Cursor] = choices[m.Cursor].GetId()
				}
			}
		case "enter":
			// do something...
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m List) View() string {
	text := ""

	for i, item := range m.Choices {
		cursor := " "
		checked := " "
		if m.Cursor == i {
			cursor = ">"
		}
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}
		text += fmt.Sprintf("%s [%s] %s\n", cursor, checked, item.GetLabel())
	}

	return text
}
