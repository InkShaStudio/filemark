package marks

import (
	"fmt"

	"github.com/InkShaStudio/filemark/pkg/storage"
	"github.com/InkShaStudio/go-command"
)

func change() *command.SCommand {

	id := command.NewCommandArg[int]("id").ChangeDescription("change mark id")
	name := command.NewCommandFlag[string]("name").ChangeDescription("change mark name")
	desc := command.NewCommandFlag[string]("desc").ChangeDescription("change mark description")
	color := command.NewCommandFlag[string]("color").ChangeDescription("change mark color")
	icon := command.NewCommandFlag[string]("icon").ChangeDescription("change mark icon")

	change := command.
		NewCommand("change").
		ChangeDescription("change mark").
		AddArgs(id).
		AddFlags(name, desc, color, icon).
		RegisterHandler(func(cmd *command.SCommand) {
			if mark, err := storage.QueryMark(id.Value); err == nil {
				if name.Value != "" {
					mark.Mark = name.Value
				}
				if desc.Value != "" {
					mark.Description = desc.Value
				}
				if color.Value != "" {
					mark.Color = color.Value
				}
				if icon.Value != "" {
					mark.Icon = icon.Value
				}
				if storage.ChangeMark(id.Value, &mark) {
					fmt.Println("change mark success")
				} else {
					fmt.Println("change mark fail")
				}
			}
		})

	return change
}
