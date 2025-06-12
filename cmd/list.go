package cmd

import (
	"os"

	"github.com/InkShaStudio/filemark/pkg/command"
)

func init() {
	p := command.NewCommandFlag[string]("path")

	rootCmd.
		AddSubCommand(
			command.
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
				}),
		)
}
