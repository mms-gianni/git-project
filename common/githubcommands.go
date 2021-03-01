package githubcommands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"gopkg.in/ukautz/clif.v1"
)

var ctx = context.Background()

func init() {
	fmt.Println("This will get called on main initialization")
}

func CreatePersonalList(c *clif.Command, in clif.Input) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Option("githubtoken").String()},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	name := ""
	if c.Argument("name") == nil {
		name = c.Argument("name").String()
	} else {
		name = in.Ask("Define the name of the new todo list: ", nil)
	}
	project, _, projectErr := client.Users.CreateProject(ctx, &github.CreateUserProjectOptions{Name: name})

	if projectErr == nil {
		body := ""
		if c.Option("description") != nil {
			body = c.Option("description").String()
		}
		public := false
		if c.Option("public").Bool() {
			public = true
		}
		client.Projects.UpdateProject(ctx, project.GetID(), &github.ProjectOptions{Body: &body, Public: &public})
		client.Projects.CreateProjectColumn(ctx, project.GetID(), &github.ProjectColumnOptions{Name: "open"})
		client.Projects.CreateProjectColumn(ctx, project.GetID(), &github.ProjectColumnOptions{Name: "closed"})
	}

}

func GetItems(c *clif.Command) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Option("githubtoken").String()},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	for _, project := range getProjects(client) {
		fmt.Println("\n\nList: ", project.GetName(), "("+project.GetState()+")")
		fmt.Println("--------------------------------------")
		cards := getCards(client, project)
		for _, card := range cards {
			fmt.Println("  <"+card.GetColumnName()+">", " ", card.GetNote())
		}
	}
}

func CreateItem(c *clif.Command, in clif.Input) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Option("githubtoken").String()},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	selectedProject := selectProject(client, in)

	projectColumns, _, _ := client.Projects.ListProjectColumns(ctx, selectedProject.GetID(), nil)
	fmt.Println(projectColumns[0].GetID(), projectColumns[0].GetName())

	message := in.Ask("What is the task", nil)
	fmt.Println(message)
	client.Projects.CreateProjectCard(ctx, projectColumns[0].GetID(), &github.ProjectCardOptions{Note: message})

}

func selectProject(client *github.Client, in clif.Input) *github.Project {
	choices := make(map[string]string)

	userprojects := getProjects(client)
	for key, project := range userprojects {
		choices[strconv.Itoa(key)] = project.GetName()
	}

	selectedNr, _ := strconv.Atoi(in.Choose("Where do you want to add a task?", choices))
	return userprojects[selectedNr]
}

func getProjects(client *github.Client) []*github.Project {

	// https://pkg.go.dev/github.com/google/go-github/v33/github#OrganizationsService.ListProjects
	// https://pkg.go.dev/github.com/google/go-github/v33/github#Project
	userprojects, res, err := client.Users.ListProjects(ctx, "mms-gianni", nil)
	//fmt.Println(userprojects)
	fmt.Println(res.Status)
	fmt.Println(err)

	return userprojects
}

func getCards(client *github.Client, project *github.Project) []*github.ProjectCard {
	var cardslist []*github.ProjectCard

	projectColumns, _, _ := client.Projects.ListProjectColumns(ctx, project.GetID(), nil)

	for _, column := range projectColumns {
		// fmt.Println(column.GetName())
		cards, _, _ := client.Projects.ListProjectCards(ctx, column.GetID(), nil)

		for _, card := range cards {
			card.ColumnName = column.Name // fix for empty card.GetColumnName()
			cardslist = append(cardslist, card)
		}
	}

	return cardslist
}
