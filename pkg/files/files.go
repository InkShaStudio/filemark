package files

import (
	"os"

	"github.com/InkShaStudio/filemark/pkg/command"
)

func Register() *command.SCommand {
	p := command.NewCommandFlag[string]("path")

	list := command.
		NewCommand("list").
		AddFlags(p).
		ChangeDescription("list all file").
		RegisterHandler(func(cmd *command.SCommand) {
			file_path := p.Value

			println("p=", file_path)

			if file_path == "" {
				file_path, _ = os.Getwd()
			}

			info, err := os.Stat(file_path)

			if err != nil {
				println("err", err)
			}

			if info.IsDir() {
				dirs, _ := os.ReadDir(file_path)
				for _, item := range dirs {
					prefix := "[F]"
					if item.IsDir() {
						prefix = "[D]"
					}
					println(prefix, item.Name())
				}
			} else {
				println("[F]", info.Name())
			}
		})

	return list
}
