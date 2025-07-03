package marks

import (
	"fmt"
	"os"

	"github.com/InkShaStudio/filemark/pkg/storage"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

type MarkList struct {
	Choices  *[]storage.Mark
	Cursor   int
	Selected map[int]int
	Multiple bool
}

func InitialListUI(list *[]storage.Mark) MarkList {
	return MarkList{
		Choices:  list,
		Cursor:   0,
		Selected: make(map[int]int),
	}
}

func (m MarkList) Run() {
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("raw mode error: ", err)
		os.Exit(1)
	}
}

func (m MarkList) Init() tea.Cmd {
	return nil
}

func (m MarkList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	choices := *m.Choices
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
					m.Selected[m.Cursor] = choices[m.Cursor].Id
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
					m.Selected[m.Cursor] = choices[m.Cursor].Id
				}
			}
		case "enter":
			// do something...
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m MarkList) View() string {
	text := ""

	for i, item := range *m.Choices {
		cursor := " "
		checked := " "
		if m.Cursor == i {
			cursor = ">"
		}
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}
		text += fmt.Sprintf("%s [%s] %s\n", cursor, checked, item.Mark)
	}

	return text
}
