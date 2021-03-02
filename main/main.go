package main

import (
	"github.com/mms-gianni/git-project/commands"
	"gopkg.in/ukautz/clif.v1"
)

func addDefaultOptions(cli *clif.Cli) {
	githubtoken := clif.NewOption("githubtoken", "t", "Private Github Token", "", true, false).
		SetEnv("GITHUB_TOKEN")
	cli.AddDefaultOptions(githubtoken)
}

func main() {
	cli := clif.New("git-project", "DEV-VERSION", "Manage your github projects with git cli")
	addDefaultOptions(cli)

	for _, cb := range commands.Commands {
		cli.Add(cb())
	}

	cli.Run()
}
