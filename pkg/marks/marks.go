package marks

import (
	"fmt"

	"github.com/InkShaStudio/filemark/pkg/storage"
	"github.com/InkShaStudio/filemark/pkg/ui"
	"github.com/InkShaStudio/go-command"
	"github.com/charmbracelet/lipgloss"
)

func Register() *command.SCommand {
	details := command.NewCommandFlag[bool]("details").ChangeDescription("show mark details").ChangeValue(false)
	raw := command.NewCommandFlag[bool]("raw").ChangeDescription("show raw content").ChangeValue(false)

	add_command := add()
	remove_command := remove()
	rename_command := rename()
	change_command := change()

	sub_commands := []*command.SCommand{
		add_command,
		remove_command,
		rename_command,
		change_command,
	}

	mark := command.
		NewCommand("mark").
		ChangeDescription("list all marks").
		AddFlags(details, raw).
		AddSubCommand(sub_commands...).
		RegisterHandler(func(cmd *command.SCommand) {
			marks := storage.QueryMarks()
			text := ""

			if raw.Value {
				model := ui.InitialListUI(
					ui.WrapItem(
						ui.WrapItemSlice(&marks),
						func(item *storage.Mark) string {
							return fmt.Sprintf("%d", item.Id)
						},
						func(item *storage.Mark) string {
							return item.Mark
						}),
					true,
				)

				model.Run()

				for i, item := range model.Choices {
					if _, ok := model.Selected[i]; ok {
						fmt.Println("u select ", item.GetLabel())
					}
				}

				operates := ui.InitialListUI(
					ui.WrapItem(
						sub_commands,
						func(item *command.SCommand) string {
							return item.Name
						},
						func(item *command.SCommand) string {
							return item.Name
						},
					),
					false,
				)

				operates.Run()

				// for i, item := range operates.Choices {
				// 	if _, ok := operates.Selected[i]; ok {
				// 		fmt.Println()
				// 	}
				// }

			} else {
				for _, mark := range marks {
					icon := ""
					color := "#16F0AE"

					if mark.Icon != "" {
						icon = mark.Icon
					}

					if mark.Color != "" {
						if mark.Color[0] != '#' {
							if new_color, err := ui.TransformColorHex(mark.Color); err == nil {
								color = new_color
							}
						} else {
							color = mark.Color
						}
					}

					fc := lipgloss.NewStyle().Foreground(lipgloss.Color(color))

					text += fc.Render(fmt.Sprintf("%s %03d %s", icon, mark.Id, mark.Mark)) + "\n"
				}
			}

			fmt.Println(text)
		})

	return mark
}
