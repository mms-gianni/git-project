package commands

import (
	gitcommands "../common"
	githubcommands "../common"

	"gopkg.in/ukautz/clif.v1"
)

func cmdCreate() *clif.Command {
	cb := func(c *clif.Command, out clif.Output, in clif.Input) {
		_, repo := gitcommands.GetGitdir()

		if repo == nil {
			githubcommands.CreatePersonalProject(c, in)
		} else {
			space := "2"

			if c.Option("profile").Bool() {
				space = "1"
			} else {
				space = in.Choose("This directory seems to be a repo. In which space do you want to create the project?", map[string]string{
					"1": "Profile",
					"2": "Repository",
				})
			}
			if space == "1" {
				githubcommands.CreatePersonalProject(c, in)
			} else {
				githubcommands.CreateRepoProject(c, in, repo)
			}
		}

	}

	return clif.NewCommand("create", "Create a new project", cb).
		NewArgument("name", "Name of the new Project", "", false, false).
		NewOption("description", "d", "Description", "", false, false).
		NewFlag("public", "p", "Make this project public (Organisations only)", false).
		NewFlag("profile", "u", "Create it in your user profile", false)
}

func init() {
	Commands = append(Commands, cmdCreate)
}
