package cmd

import (
	"fmt"
	"os"

	"github.com/InkShaStudio/filemark/pkg/command"
	"github.com/InkShaStudio/filemark/pkg/storage"
)

var rootCmd = command.
	NewCommand("filemark").
	ChangeDescription("filemark is a tool\ncan add tag to file or dir").
	RegisterHandler(func(cmd *command.SCommand) {
	})

func init() {
	storage.CreateTable()
}

func Execute() {
	cmd := command.RegisterCommand(rootCmd)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// data, err := json.Marshal(rootCmd)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// } else {
	// 	os.WriteFile("cmd.json", data, 0644)
	// }
}
