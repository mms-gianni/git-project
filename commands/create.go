package commands

import (
	githubcommands "github.com/mms-gianni/git-project/common"

	"gopkg.in/ukautz/clif.v1"
)

func cmdCreate() *clif.Command {
	cb := func(c *clif.Command, out clif.Output, in clif.Input) {
		githubcommands.CreateProject(c, in)
	}

	return clif.NewCommand("create", "Create a new project", cb).
		NewArgument("project", "Name of the new Project", "", false, false).
		NewOption("description", "d", "Description", "", false, false).
		NewFlag("public", "p", "Make this project public (Organisations only)", false).
		NewFlag("profile", "u", "Create it in your user profile", false)
}

func init() {
	Commands = append(Commands, cmdCreate)
}
