package main

import (
	"os"

	"github.com/mms-gianni/git-project/commands"
	"gopkg.in/ukautz/clif.v1"
)

func addDefaultOptions(cli *clif.Cli) {
	githubtoken := clif.NewOption("githubtoken", "t", "Private Github Token", "", true, false).
		SetEnv("GITHUB_TOKEN")

	githubusername := clif.NewOption("username", "u", "Github username", "", false, false).
		SetEnv("GITHUB_USERNAME")

	githubOrganisations := clif.NewOption("organisations", "o", "Github organisations (comma separated)", "", false, false).
		SetEnv("GITHUB_ORGANISATIONS")
	cli.AddDefaultOptions(githubtoken, githubusername, githubOrganisations)
}

func main() {
	cli := clif.New("git-project", "DEV-VERSION", "Manage your github projects with git cli")

	var OwnStyles = map[string]string{
		"error":       "\033[31;1m",
		"warn":        "\033[33m",
		"info":        "\033[0;97m",
		"success":     "\033[32m",
		"debug":       "\033[30;1m",
		"headline":    "\033[4;1m",
		"subline":     "\033[4m",
		"important":   "\033[47;30;1m",
		"query":       "\033[36m",
		"reset":       "\033[0m",
		"new":         "\U00002728",
		"open":        "\U00002B50",
		"done":        "\U00002705",
		"closed":      "\U0001F512",
		"in progress": "\U0001F528",
		"waiting":     "\U0001F551",
		"critical":    "\U000026A0",
		"idea":        "\U0001F4A1",
	}

	cli.SetOutput(clif.NewColorOutput(os.Stdout).SetFormatter(clif.NewDefaultFormatter(OwnStyles)))

	addDefaultOptions(cli)

	for _, cb := range commands.Commands {
		cli.Add(cb())
	}

	cli.Run()
}

/*
H:git-projectR: D:(DEV-VERSION)R:
I:Manage your github projects with git cliR:

U:Usage:R:
  main command [arg ..] [--opt val ..]

U:Available commands:R:
  I:add   R:  Add a new card
  I:clean R:  Archive all cards in the 'done' column
  I:close R:  Close a project
  I:createR:  Create a new project
  I:help  R:  Show this help
  I:list  R:  List all available commands
  I:statusR:  List all projects and cards
*/
