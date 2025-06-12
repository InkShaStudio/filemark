package cmd

import (
	"fmt"

	"github.com/InkShaStudio/filemark/pkg/command"
	"github.com/InkShaStudio/filemark/pkg/storage"
)

func init() {
	name := command.NewCommandArg[string]("name").ChangeDescription("mark name")
	color := command.NewCommandArg[string]("color").ChangeDescription("mark color")
	description := command.NewCommandArg[string]("description").ChangeDescription("mark description")
	icon := command.NewCommandArg[string]("icon").ChangeDescription("mark icon")

	add := command.
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

	rootCmd.
		AddSubCommand(
			command.
				NewCommand("mark").
				ChangeDescription("list all marks").
				AddSubCommand(add).
				RegisterHandler(func(cmd *command.SCommand) {
					marks := storage.QueryMarks()
					for index, mark := range marks {
						fmt.Printf("[%d] %s\n", index, mark.Mark)
					}
				}),
		)
}
