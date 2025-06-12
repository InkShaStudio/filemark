package marks

import (
	"github.com/InkShaStudio/go-command"
)

func rename() *command.SCommand {
	rename := command.
		NewCommand("rename")

	return rename
}
