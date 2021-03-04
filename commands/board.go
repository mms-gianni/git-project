package commands

import (
	githubcommands "github.com/mms-gianni/git-project/common"
	"gopkg.in/ukautz/clif.v1"
)

func cmdBoard() *clif.Command {
	cb := func(c *clif.Command, out clif.Output) {
		githubcommands.GetBoard(c, out)
	}

	return clif.NewCommand("board", "Display project boards", cb).
		NewArgument("project", "Show only this project", "", false, false)
}

func init() {
	Commands = append(Commands, cmdBoard)
}
