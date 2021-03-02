package commands

import (
	githubcommands "github.com/mms-gianni/git-project/common"
	"gopkg.in/ukautz/clif.v1"
)

func cmdAdd() *clif.Command {
	cb := func(c *clif.Command, out clif.Output, in clif.Input) {
		githubcommands.CreateItem(c, in)
	}

	return clif.NewCommand("add", "Add a new card", cb).
		NewArgument("project", "Show only repo", "", false, false).
		NewArgument("message", "Message", "<empty>", false, false).
		NewArgument("status", "Status", "open", false, false)
}

func init() {
	Commands = append(Commands, cmdAdd)
}
