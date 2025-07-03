package files

import (
	"os"

	"github.com/InkShaStudio/go-command"
)

func Register() *command.SCommand {
	p := command.NewCommandFlag[string]("path").ChangeDescription("work dir path")

	list := command.
		NewCommand("list").
		AddFlags(p).
		ChangeDescription("list all file").
		RegisterHandler(func(cmd *command.SCommand) {
			file_path := p.Value

			if file_path == "" {
				file_path, _ = os.Getwd()
			}

			filter := NewFileInfoFilter()
			list := ReadPath(file_path)
			ui := NewFileInfoList(filter, &list)
			ui.CurentPath = file_path

			println("ready show ui")

			ui.Run()
		})

	return list
}
