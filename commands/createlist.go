package commands

import (
	"fmt"
	"log"
	"os"
	"strings"

	githubcommands "../common"

	"github.com/go-git/go-git"
	"gopkg.in/ukautz/clif.v1"
)

func cmdCreatelist() *clif.Command {
	cb := func(c *clif.Command, out clif.Output, in clif.Input) {
		out.Printf("Create a new list\n")

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		alldirs := strings.Split(dir, "/")
		fmt.Println(alldirs)

		testdir := ""
		for _, dirname := range alldirs[1:] {

			testdir = testdir + "/" + dirname
			repo, giterror := git.PlainOpen(testdir)

			if giterror != nil {
				fmt.Println(giterror, "in", testdir)
			} else {
				fmt.Println(testdir, "is a git dir")
				fmt.Println(repo.Remotes())
				break
			}

		}

		githubcommands.CreatePersonalList(c, in)
	}

	return clif.NewCommand("createlist", "Add a new todo list", cb).
		NewArgument("name", "Name of the new Project", "", false, false).
		NewArgument("repo", "create in repo", "", false, false).
		NewOption("description", "d", "Description", "", false, false).
		NewFlag("public", "p", "Make this todolist public", false)
}

func init() {
	Commands = append(Commands, cmdCreatelist)
}
