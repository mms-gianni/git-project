package githubcommands

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GetProject(token string) {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	userprojects, res, err := client.Users.ListProjects(ctx, "mms-gianni", nil)
	fmt.Println(userprojects)
	fmt.Println(res)
	fmt.Println(err)
	fmt.Println(token)
}
