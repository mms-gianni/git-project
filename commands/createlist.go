package commands

import (
	gitcommands "../common"
	githubcommands "../common"

	"gopkg.in/ukautz/clif.v1"
)

func cmdCreate() *clif.Command {
	cb := func(c *clif.Command, out clif.Output, in clif.Input) {
		out.Printf("Create a new list\n")

		_, repo := gitcommands.GetGitdir()

		if repo == nil {
			githubcommands.CreatePersonalList(c, in)
		} else {
			space := in.Choose("This directory seems to be a repo. In which space do you want to create the repo?", map[string]string{
				"1": "Personal",
				"2": "Repo",
			})
			if space == "1" {
				githubcommands.CreatePersonalList(c, in)
			} else {
				githubcommands.CreateRepoProject(c, in, repo)
			}
		}

	}

	return clif.NewCommand("create", "Create a new Project", cb).
		NewArgument("name", "Name of the new Project", "", false, false).
		NewArgument("repo", "create in repo", "", false, false).
		NewOption("description", "d", "Description", "", false, false).
		NewFlag("public", "p", "Make this todolist public", false)
}

func init() {
	Commands = append(Commands, cmdCreate)
}
