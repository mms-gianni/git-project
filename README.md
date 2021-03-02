# git-project
Manage your github projects with your git cli

## Why
- atomate your projects
- manage your projects where you work
- use it as your personal todo list

## Installation
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

### close a obsolete project (can be reopened on github)
```
git project close
```
## help and shortcuts
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
