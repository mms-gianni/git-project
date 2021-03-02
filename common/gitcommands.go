package common

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
)

type repoDetails struct {
	name  string
	owner string
}

func GetGitdir() (gitBasedir *string, repo *git.Repository) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	alldirs := strings.Split(dir, "/")
	//fmt.Println(alldirs)

	testdir := ""
	for _, dirname := range alldirs[1:] {

		testdir = testdir + "/" + dirname
		repo, giterror := git.PlainOpen(testdir)

		if giterror != nil {
			//fmt.Println(giterror, "in", testdir)
		} else {
			fmt.Println(testdir, "is a git dir")
			return &testdir, repo
		}

	}
	return nil, nil
}

func getRepodetails(repo *git.Repository) (r *repoDetails) {
	remotes, _ := repo.Remotes()
	re := regexp.MustCompile(`.*git@github.com:(.*)/(.*)\.git \(fetch\)`)
	findings := re.FindAllStringSubmatch(remotes[0].String(), -1)

	repodetails := repoDetails{owner: findings[0][1], name: findings[0][2]}

	return &repodetails
}
