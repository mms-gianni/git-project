package commands

import (
	githubcommands "github.com/mms-gianni/git-project/common"
	"gopkg.in/ukautz/clif.v1"
)

func cmdStatus() *clif.Command {
	cb := func(c *clif.Command, out clif.Output) {
		out.Printf("a long list.\n")
		githubcommands.GetItems(c)
	}

	return clif.NewCommand("status", "List all projects and cards", cb).
		NewArgument("repo", "Show only repo", "", false, false)
}

func init() {
	Commands = append(Commands, cmdStatus)
}
