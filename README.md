# git-project
Manage your github projects with your git cli

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mms-gianni/git-project)
![GitHub top language](https://img.shields.io/github/languages/top/mms-gianni/git-project)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/mms-gianni/git-project/Upload%20Release%20Asset)
## Why
- atomate your projects
- manage your projects where you work
- use it as your personal todo list

## Installation
Generate a token here : https://github.com/settings/tokens (You need to be loged in)
```
export GITHUB_TOKEN="asdfasdfasdfasdfasdfasdfasdfasdfasdf"
curl https://raw.githubusercontent.com/mms-gianni/git-project/master/cmd/git-project/git-project.mac.64bit -o /usr/local/bin/git-project
```


## Quick start

### Create your first personl project in your profile
```
git project create -u
```

### Create a repository related project
```
cd your-project
git project create 
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

### Cleanup all Cards in state "closed"
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
git project help create
```

### create a personal list with one shot
```
git project create shoppinglist -p -d "helps me to remember what to buy"
```
