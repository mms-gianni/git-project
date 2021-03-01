package common

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-git/go-git"
)

func GetGitdir() (gitBasedir *string, repo *git.Repository) {
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
			return &testdir, repo
		}

	}
	return nil, nil
}
