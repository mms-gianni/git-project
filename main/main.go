package main

import (
	"../commands"
	"gopkg.in/ukautz/clif.v1"
)

func addDefaultOptions(cli *clif.Cli) {
	githubtoken := clif.NewOption("githubtoken", "t", "Private Github Token", "", true, false).
		SetEnv("GITHUB_TOKEN")
	cli.AddDefaultOptions(githubtoken)
}

func main() {
	cli := clif.New("git-todo", "0.0.1", "An cli todo list connected with github")
	addDefaultOptions(cli)

	for _, cb := range commands.Commands {
		cli.Add(cb())
	}

	cli.Run()
}
