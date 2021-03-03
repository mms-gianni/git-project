package commands

import (
	githubcommands "github.com/mms-gianni/git-project/common"
	"gopkg.in/ukautz/clif.v1"
)

func cmdClean() *clif.Command {
	cb := func(c *clif.Command, out clif.Output, in clif.Input) {
		githubcommands.Cleanup(c, out)
	}

	return clif.NewCommand("clean", "Archive all cards in the 'done' column", cb)
}

func init() {
	Commands = append(Commands, cmdClean)
}
