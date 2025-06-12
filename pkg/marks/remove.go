package marks

import (
	"fmt"

	"github.com/InkShaStudio/go-command"
	"github.com/InkShaStudio/filemark/pkg/storage"
)

func remove() *command.SCommand {
	id := command.NewCommandArg[int]("id").ChangeDescription("remove mark id")

	return command.
		NewCommand("remove").
		ChangeDescription("remove mark <id>").
		AddArgs(id).
		RegisterHandler(func(cmd *command.SCommand) {
			if storage.RemoveMark(id.Value) {
				fmt.Println("remove mark success")
			} else {
				fmt.Println("remove mark fail")
			}
		})
}
