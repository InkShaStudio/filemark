package marks

import (
	"fmt"

	"github.com/InkShaStudio/filemark/pkg/command"
	"github.com/InkShaStudio/filemark/pkg/storage"
)

func Register() *command.SCommand {

	mark := command.
		NewCommand("mark").
		ChangeDescription("list all marks").
		AddSubCommand(add(), remove()).
		RegisterHandler(func(cmd *command.SCommand) {
			marks := storage.QueryMarks()
			for _, mark := range marks {
				fmt.Printf("[%d] %s\n", mark.Id, mark.Mark)
			}
		})

	return mark
}
