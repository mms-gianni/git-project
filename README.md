# git-project
Manage your github projects with your git cli

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mms-gianni/git-project)
![GitHub top language](https://img.shields.io/github/languages/top/mms-gianni/git-project)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/mms-gianni/git-project/Upload%20Release%20Asset)
![GitHub MIT license](https://img.shields.io/github/license/mms-gianni/git-project)
![Swiss made](https://img.shields.io/badge/swiss%20made-100%25-red)
## Why
- atomate your projects
- manage your projects where you work
- use it as your personal todo list

## Installation
Generate a token here : https://github.com/settings/tokens (You need to be loged in)

To export the Github username is optional. But it requred to access your personal profile projects. 
### Mac
```
echo 'export GITHUB_TOKEN="asdfasdfasdfasdfasdfasdfasdfasdfasdf"' >> ~/.zshrc
echo 'export GITHUB_USERNAME="change-me-to-your-username"' >> ~/.zshrc
curl https://raw.githubusercontent.com/mms-gianni/git-project/master/cmd/git-project/git-project.mac.64bit -o /usr/local/bin/git-project
chmod +x /usr/local/bin/git-project
```

### Linux 
```
echo 'export GITHUB_TOKEN="asdfasdfasdfasdfasdfasdfasdfasdfasdf"' >> ~/.bashrc
echo 'export GITHUB_USERNAME="change-me-to-your-username"' >> ~/.bashrc
curl https://raw.githubusercontent.com/mms-gianni/git-project/master/cmd/git-project/git-project.linux.64bit -o /usr/local/bin/git-project
chmod +x /usr/local/bin/git-project
```

### Windows
It works also on Windows. But since i have no windows in my basement, I have no idea how to install it (Help welcome :D )

You find older releases here : https://github.com/mms-gianni/git-project/releases

## Quick start

### Create your first personl project in your profile
Open creates a new project. 2 columns are added by default (open, done). Feel free to change them or add more.
```
git project open -U
```

### Create a repository related project
```
cd your-project
git project open 
```

### Add a new task to a project
```
git project add
```

### Show a overview of your projects
(note repository related projects will only be displayed when you are in the current workindir)

```
git project status
```

### Display project boards

```
git project board
```

### Move a card to another column
```
git project move
```

### Cleanup all Cards in column "done"
```
git project clean
```

### Close a obsolete project (can be reopened on github)
```
git project close
```
## Help and shortcuts
```
git project list
```

```
git project help open
```

### Create a personal list with one shot
```
git project create shoppinglist -p -d "helps me to remember what to buy"
```

### Add Milk to your shoppinglist
```
git project add shoppinglist milk
```

### Move a card arround 
```
git project move shoppinglist -c milk -d done
```