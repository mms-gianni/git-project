package commands

import (
	githubcommands "github.com/mms-gianni/git-project/common"
	"gopkg.in/ukautz/clif.v1"
)

func cmdClose() *clif.Command {
	cb := func(c *clif.Command, out clif.Output, in clif.Input) {
		githubcommands.CloseProject(c, in)
	}

	return clif.NewCommand("close", "Close a project", cb).
		NewArgument("project", "Name of the project to close", "", false, false)
}

func init() {
	Commands = append(Commands, cmdClose)
}
