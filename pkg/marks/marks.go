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

	return command.
		NewCommand("mark").
		ChangeDescription("list all marks").
		AddFlags(details, raw).
		AddSubCommand(sub_commands...).
		RegisterHandler(func(cmd *command.SCommand) {
			marks := storage.QueryMarks()
			text := ""

			for _, mark := range marks {
				icon := ui.DEFAULT_MARK_ICON
				color := ui.DEFAULT_MARK_COLOR

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

			fmt.Println(text)
		})
}
