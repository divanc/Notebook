

# How To Github
Okay, that may sound strange, but making a durable pauses I often forget how to do some basic tasks via Git.  
First of all, how to write markdown texts like this one. I guess [Github Tutorial](https://guides.github.com/features/mastering-markdown/) is nice enough, but let's talk commands:


### Practical stuff

```console
divan@giter:~$ git init                  # Creates a NEW REPO IN CURRENT FOLDER

divan@giter:~$ git clone *url*           # Downloads from url

divan@giter:~$ git add filename          # Sends file to the query

divan@giter:~$ git commit -m "What A Lovely Day" #Commits, dah?

divan@giter:~$ git push                  # Sends to Hub

divan@giter:~$ git pull                  # Takes updates from server to local

divan@giter:~$ git checkout *name*       # Goes to branch      (and then git pull)

divan@giter:~$ git checkout -b *name*    # Creates a new branch and goes to it

divan@giter:~$ git branch *name*         # Just creates branch

divan@giter:~$ git merge *branch*        # Merges *branch* to current branch

divan@giter:~$ git request-pull



```

### Extra stuff
```console
divan@giter:~$ git status          # Checks wether there are ambiguos ambivalent stuff

divan@giter:~$ git checkout        # Cancels & drops changes

divan@giter:~$ git reset           # Cancels on 'add' stage

divan@giter:~$ git log --pretty=oneline # dah

divan@giter:~$ git revert HEAD     # Cancels last commit
```

### Tags
```console
divan@giter:~$ git tag -a v1.0

divan@giter:~$ git push -tags

divan@giter:~$ git checkout *tag*

```


### First Config

```console
divan@giter:~$ git config --global user.name "Divan"
divan@giter:~$ git config --global user.email "SendMeA@mail.com"

divan@giter:~$ git config --global core.autocrlf input
divan@giter:~$ config --global core.safecrlf true 

divan@giter:~$ git config --global core.quotepath off
```
