package cmd

import (
	"fmt"
	"os"

	"github.com/InkShaStudio/filemark/pkg/command"
	"github.com/InkShaStudio/filemark/pkg/files"
	"github.com/InkShaStudio/filemark/pkg/marks"
	"github.com/InkShaStudio/filemark/pkg/storage"
)

func init() {
	storage.CreateTable()
}

func Execute() {
	cmd := command.RegisterCommand(
		command.
			NewCommand("filemark").
			ChangeDescription("filemark is a tool\ncan add tag to file or dir").
			RegisterHandler(func(cmd *command.SCommand) {
			}).
			AddSubCommand(
				files.Register(),
				marks.Register(),
			),
	)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
