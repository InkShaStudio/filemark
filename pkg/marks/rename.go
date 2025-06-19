package marks

import (
	"fmt"

	"github.com/InkShaStudio/filemark/pkg/storage"
	"github.com/InkShaStudio/go-command"
)

func rename() *command.SCommand {
	id := command.NewCommandArg[int]("id").ChangeDescription("change mark id")
	name := command.NewCommandArg[string]("name").ChangeDescription("new mark name")

	rename := command.
		NewCommand("rename").
		ChangeDescription("rename mark").
		AddArgs(id, name).
		RegisterHandler(func(cmd *command.SCommand) {
			if storage.RenameMark(id.Value, name.Value) {
				fmt.Println("rename mark success")
			} else {
				fmt.Println("rename mark fail")
			}
		})

	return rename
}
