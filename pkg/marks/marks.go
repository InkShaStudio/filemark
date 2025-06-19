package marks

import (
	"fmt"
	"os"

	"github.com/InkShaStudio/filemark/pkg/storage"
	"github.com/InkShaStudio/filemark/pkg/ui"
	"github.com/InkShaStudio/go-command"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func Register() *command.SCommand {
	details := command.NewCommandFlag[bool]("details").ChangeDescription("show mark details").ChangeValue(false)
	raw := command.NewCommandFlag[bool]("raw").ChangeDescription("show raw content").ChangeValue(false)

	add_command := add()
	remove_command := remove()
	rename_command := rename()
	change_command := change()

	mark := command.
		NewCommand("mark").
		ChangeDescription("list all marks").
		AddFlags(details, raw).
		AddSubCommand(
			add_command,
			remove_command,
			rename_command,
			change_command,
		).
		RegisterHandler(func(cmd *command.SCommand) {
			marks := storage.QueryMarks()
			text := ""
			if raw.Value {
				model := ui.InitialListUI(
					ui.WrapItem(
						&marks,
						func(item *storage.Mark) string {
							return fmt.Sprintf("%d", item.Id)
						},
						func(item *storage.Mark) string {
							return item.Mark
						}),
				)

				if _, err := tea.NewProgram(model).Run(); err != nil {
					fmt.Println("raw mode error: ", err)
					os.Exit(1)
				}

				for i, item := range model.Choices {
					if _, ok := model.Selected[i]; ok {
						fmt.Println("u select ", item.GetLabel())
					}
				}

			} else {
				for _, mark := range marks {
					icon := ""
					color := "#16F0AE"

					if mark.Icon != "" {
						icon = mark.Icon
					}
					if mark.Color != "" {
						color = mark.Color
					}
					fc := lipgloss.NewStyle().Foreground(lipgloss.Color(color))

					text += fc.Render(fmt.Sprintf("%s %03d %s", icon, mark.Id, mark.Mark)) + "\n"
				}
			}

			fmt.Println(text)
		})

	return mark
}
