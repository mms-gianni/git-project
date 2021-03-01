package commands

import (
	githubcommands "../common"
	"gopkg.in/ukautz/clif.v1"
)

func cmdList() *clif.Command {
	cb := func(c *clif.Command, out clif.Output) {
		out.Printf("a long list.\n")
		githubcommands.GetItems(c)
	}

	return clif.NewCommand("status", "List all Todo", cb).
		NewArgument("repo", "Show only repo", "", false, false)
}

func init() {
	Commands = append(Commands, cmdList)
}
