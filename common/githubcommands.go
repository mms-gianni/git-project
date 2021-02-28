package githubcommands

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GetItems(token string) {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// https://pkg.go.dev/github.com/google/go-github/v33/github#OrganizationsService.ListProjects
	userprojects, res, err := client.Users.ListProjects(ctx, "mms-gianni", nil)
	//fmt.Println(userprojects)
	fmt.Println(res.Status)
	fmt.Println(err)

	// https://pkg.go.dev/github.com/google/go-github/v33/github#Project
	for _, project := range userprojects {
		//fmt.Println(project.GetID())
		//fmt.Println(project.GetHTMLURL())
		fmt.Println("List: ", project.GetName(), "("+project.GetState()+")")
		fmt.Println("_______________________________________")

		projectColumns, _, _ := client.Projects.ListProjectColumns(ctx, project.GetID(), nil)

		for _, column := range projectColumns {
			// fmt.Println(column.GetName())
			cards, _, _ := client.Projects.ListProjectCards(ctx, column.GetID(), nil)

			for _, card := range cards {
				fmt.Println(column.GetName(), " ", card.GetNote())
			}
		}
	}
}
