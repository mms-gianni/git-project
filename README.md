# git-project
Manage your github projects with yout git cli

## Why
- atomate your projects
- use it as your personal todo list

## installation
```
curl https://raw.githubusercontent.com/mms-gianni/git-project/master/cms/git-project.mac.64bit -o /usr/local/bin/git-project
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