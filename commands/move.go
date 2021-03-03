package commands

import (
	githubcommands "github.com/mms-gianni/git-project/common"
	"gopkg.in/ukautz/clif.v1"
)

func cmdMove() *clif.Command {
	cb := func(c *clif.Command, out clif.Output, in clif.Input) {
		githubcommands.MoveCard(c, out, in)

	}

	return clif.NewCommand("move", "Move a card to anoter column", cb).
		NewArgument("project", "Name of the project", "", false, false).
		NewOption("card", "c", "Selected cards", "", false, false).
		NewOption("destinateion", "d", "destination column", "closed", false, false)
}

func init() {
	Commands = append(Commands, cmdMove)
}
