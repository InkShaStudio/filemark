package marks

import (
	"fmt"

	"github.com/InkShaStudio/filemark/pkg/command"
	"github.com/InkShaStudio/filemark/pkg/storage"
)

func add() *command.SCommand {
	name := command.NewCommandArg[string]("name").ChangeDescription("mark name")
	color := command.NewCommandArg[string]("color").ChangeDescription("mark color")
	description := command.NewCommandArg[string]("description").ChangeDescription("mark description")
	icon := command.NewCommandArg[string]("icon").ChangeDescription("mark icon")

	return command.
		NewCommand("add").
		ChangeDescription("add mark").
		AddArgs(name, color, description, icon).
		RegisterHandler(func(cmd *command.SCommand) {
			storage.CreateTable()
			flag := storage.InsertMark(storage.CreateMark{
				Mark:        name.Value,
				Description: description.Value,
				Color:       color.Value,
				Icon:        icon.Value,
			})
			if flag {
				fmt.Println("add mark success")
			} else {
				fmt.Println("add mark failed")
			}
		})
}
