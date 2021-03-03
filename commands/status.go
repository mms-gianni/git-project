package commands

import (
	githubcommands "github.com/mms-gianni/git-project/common"
	"gopkg.in/ukautz/clif.v1"
)

func cmdStatus() *clif.Command {
	cb := func(c *clif.Command, out clif.Output) {
		githubcommands.GetStatus(c, out)
	}

	return clif.NewCommand("status", "List projects and cards", cb).
		NewArgument("project", "Show only this project", "", false, false)
}

func init() {
	Commands = append(Commands, cmdStatus)
}
