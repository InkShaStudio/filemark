package marks

import (
	"fmt"

	"github.com/InkShaStudio/filemark/pkg/storage"
	"github.com/InkShaStudio/go-command"
)

func add() *command.SCommand {
	name := command.NewCommandArg[string]("name").ChangeDescription("mark name")
	description := command.NewCommandArg[string]("description").ChangeDescription("mark description").ChangeValue("")
	color := command.NewCommandFlag[string]("color").ChangeDescription("mark color").ChangeValue("#FFFFFF")
	icon := command.NewCommandFlag[string]("icon").ChangeDescription("mark icon").ChangeValue("")

	return command.
		NewCommand("add").
		ChangeDescription("create mark").
		AddArgs(name, description).
		AddFlags(color, icon).
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
