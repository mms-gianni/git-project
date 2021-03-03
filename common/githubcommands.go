package common

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
	"gopkg.in/ukautz/clif.v1"
)

var ctx = context.Background()

func login(c *clif.Command) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Option("githubtoken").String()},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}

type CardslistItem struct {
	id          int
	carddetails *github.ProjectCard
	project     *github.Project
}

func Cleanup(c *clif.Command, out clif.Output) {
	client := login(c)

	for _, project := range getProjects(client) {

		cards := getCards(client, project)
		for _, card := range cards {
			if card.GetColumnName() == "done" {
				fmt.Println("Archived", card.GetNote(), "in", project.GetName())
				archived := true
				_, _, err := client.Projects.UpdateProjectCard(ctx, card.GetID(), &github.ProjectCardOptions{Archived: &archived})
				if err == nil {
					out.Printf("\n<success> Archived <" + card.GetNote() + "> in <" + project.GetName() + "> project<reset>\n\n")
				}
			}
		}

	}
}

func CloseProject(c *clif.Command, in clif.Input, out clif.Output) {
	client := login(c)

	selectedProject := selectProject(client, in, c.Argument("project").String())

	state := "closed"
	project, _, err := client.Projects.UpdateProject(ctx, selectedProject.GetID(), &github.ProjectOptions{State: &state})

	if err == nil {
		out.Printf("\n<success> project <" + project.GetName() + "> has been sucessfully closed<reset>\n\n")
	}
}

func OpenProject(c *clif.Command, in clif.Input, out clif.Output) {
	_, repo := GetGitdir()

	if repo == nil {
		OpenPersonalProject(c, in, out)
	} else {
		space := "2"

		if c.Option("profile").Bool() {
			space = "1"
		} else {
			repodetails := getRepodetails(repo)
			space = in.Choose("This directory seems to be a repo. In which space do you want to create the project?", map[string]string{
				"1": "Profile",
				"2": "Repository (" + repodetails.name + ")",
			})
		}
		if space == "1" {
			OpenPersonalProject(c, in, out)
		} else {
			OpenRepoProject(c, in, out, repo)
		}
	}
}

func OpenRepoProject(c *clif.Command, in clif.Input, out clif.Output, repo *git.Repository) {
	client := login(c)

	repositorydetails := getRepodetails(repo)

	name := ""
	if c.Argument("project").String() != "" {
		name = c.Argument("project").String()
	} else {
		name = in.Ask("Define the name of the new project: ", nil)
	}
	body := ""
	if c.Option("description").String() != "" {
		body = c.Option("description").String()
	}
	public := false
	if c.Option("public").Bool() {
		public = true
	}
	project, _, projectErr := client.Repositories.CreateProject(ctx, repositorydetails.owner, repositorydetails.name, &github.ProjectOptions{Name: &name, Body: &body, Public: &public})
	if projectErr == nil {
		_, _, openColumnErr := client.Projects.CreateProjectColumn(ctx, project.GetID(), &github.ProjectColumnOptions{Name: "open"})
		_, _, doneColumnErr := client.Projects.CreateProjectColumn(ctx, project.GetID(), &github.ProjectColumnOptions{Name: "done"})
		if projectErr == nil && openColumnErr == nil && doneColumnErr == nil {
			out.Printf("\n<success> project <" + project.GetName() + "> has been sucessfully opened<reset>\n\n")
		}
	}
}

func OpenPersonalProject(c *clif.Command, in clif.Input, out clif.Output) {
	client := login(c)

	name := ""
	if c.Argument("project").String() != "" {
		name = c.Argument("project").String()
	} else {
		name = in.Ask("Define the name of the new project: ", nil)
	}
	project, _, projectErr := client.Users.CreateProject(ctx, &github.CreateUserProjectOptions{Name: name})

	if projectErr == nil {
		body := ""
		if c.Option("description").String() != "" {
			body = c.Option("description").String()
		}
		public := false
		if c.Option("public").Bool() {
			public = true
		}
		client.Projects.UpdateProject(ctx, project.GetID(), &github.ProjectOptions{Body: &body, Public: &public})
		_, _, openColumnErr := client.Projects.CreateProjectColumn(ctx, project.GetID(), &github.ProjectColumnOptions{Name: "open"})
		_, _, doneColumnErr := client.Projects.CreateProjectColumn(ctx, project.GetID(), &github.ProjectColumnOptions{Name: "done"})
		if projectErr == nil && openColumnErr == nil && doneColumnErr == nil {
			out.Printf("\n<success> project <" + project.GetName() + "> has been sucessfully opened<reset>\n\n")
		}
	}

}

func GetStatus(c *clif.Command, out clif.Output) []CardslistItem {
	client := login(c)

	var projectslist []*github.Project
	if c.Argument("project").String() == "" {
		projectslist = getProjects(client)
	} else {
		projectslist = append(projectslist, getProjectByName(client, c.Argument("project").String()))
	}

	item := 0
	var cardslist []CardslistItem
	for _, project := range projectslist {
		cards := getCards(client, project)
		out.Printf("\n<subline>List: " + project.GetName() + "<reset>\n")
		for _, card := range cards {
			out.Printf(strconv.Itoa(item) + "|  <" + card.GetColumnName() + ">  " + card.GetNote() + "\n")
			cardslist = append(cardslist, CardslistItem{
				id:          item,
				carddetails: card,
				project:     project,
			})
			item++
		}
	}

	return cardslist
}

func MoveCard(c *clif.Command, out clif.Output, in clif.Input) {
	client := login(c)

	selectedProject := selectProject(client, in, c.Argument("project").String())

	var selectedCard *github.ProjectCard
	cards := getCards(client, selectedProject)
	if c.Option("card").String() == "" {
		selectedCard = selectCard(cards, in)
	} else {
		selectedCard = selectCardByNote(cards, c.Option("card").String())
	}

	var selectedColumn *github.ProjectColumn
	selectedColumn = selectColumn(client, in, selectedProject, c.Option("destination").String())

	_, err := client.Projects.MoveProjectCard(ctx, selectedCard.GetID(), &github.ProjectCardMoveOptions{Position: "bottom", ColumnID: selectedColumn.GetID()})

	if err == nil {
		out.Printf("\n\nMoved '" + selectedCard.GetNote() + "' to <" + selectedColumn.GetName() + "> " + selectedColumn.GetName() + "\n")
	}

}

func CreateCard(c *clif.Command, in clif.Input, out clif.Output) {
	client := login(c)

	selectedProject := selectProject(client, in, c.Argument("project").String())

	projectColumns, _, _ := client.Projects.ListProjectColumns(ctx, selectedProject.GetID(), nil)

	message := in.Ask("What is the task", nil)

	card, _, cardErr := client.Projects.CreateProjectCard(ctx, projectColumns[0].GetID(), &github.ProjectCardOptions{Note: message})

	if cardErr == nil {
		out.Printf("\n<success> added Card <" + card.GetNote() + "> sucessfully to <" + selectedProject.GetName() + "> project<reset>\n\n")
	}

}
func selectColumn(client *github.Client, in clif.Input, project *github.Project, searchColumn string) *github.ProjectColumn {
	choices := make(map[string]string)

	columns, _, _ := client.Projects.ListProjectColumns(ctx, project.GetID(), nil)
	for key, column := range columns {
		choices[strconv.Itoa(key)] = "<" + column.GetName() + ">"
		if column.GetName() == searchColumn {
			return column
		}
	}

	selectedNr, _ := strconv.Atoi(in.Choose("Select column to move the card", choices))
	return columns[selectedNr]
}

func selectCard(cards []*github.ProjectCard, in clif.Input) *github.ProjectCard {
	choices := make(map[string]string)

	for key, card := range cards {
		choices[strconv.Itoa(key)] = "<" + card.GetColumnName() + "> " + card.GetNote()
	}

	selectedNr, _ := strconv.Atoi(in.Choose("Select Card to move", choices))
	return cards[selectedNr]
}

func selectCardByNote(cards []*github.ProjectCard, searchedCard string) *github.ProjectCard {
	for _, card := range cards {
		if card.GetNote() == searchedCard {
			return card
		}
	}
	return nil
}

func selectProject(client *github.Client, in clif.Input, preselectedProject string) *github.Project {
	choices := make(map[string]string)

	userprojects := getProjects(client)
	for key, project := range userprojects {
		choices[strconv.Itoa(key)] = project.GetName()
		if project.GetName() == preselectedProject {
			return project
		}
	}

	selectedNr, _ := strconv.Atoi(in.Choose("Where do you want to add a task?", choices))
	return userprojects[selectedNr]
}

func getProjectByName(client *github.Client, projectname string) *github.Project {
	userprojects := getProjects(client)
	for _, project := range userprojects {
		if project.GetName() == projectname {
			return project
		}
	}
	return nil
}

func getProjects(client *github.Client) []*github.Project {

	// https://pkg.go.dev/github.com/google/go-github/v33/github#OrganizationsService.ListProjects
	// https://pkg.go.dev/github.com/google/go-github/v33/github#Project
	userprojects, _, _ := client.Users.ListProjects(ctx, "mms-gianni", nil)

	_, repo := GetGitdir()

	if repo != nil {
		repositorydetails := getRepodetails(repo)

		repoprojects, _, _ := client.Repositories.ListProjects(ctx, repositorydetails.owner, repositorydetails.name, nil)

		userprojects = append(userprojects, repoprojects...)
	}

	return userprojects
}

func getCards(client *github.Client, project *github.Project) []*github.ProjectCard {
	var cardslist []*github.ProjectCard

	projectColumns, _, _ := client.Projects.ListProjectColumns(ctx, project.GetID(), nil)

	for _, column := range projectColumns {
		cards, _, _ := client.Projects.ListProjectCards(ctx, column.GetID(), nil)

		for _, card := range cards {
			card.ColumnName = column.Name // fix for empty card.GetColumnName()
			cardslist = append(cardslist, card)
		}
	}

	return cardslist
}
