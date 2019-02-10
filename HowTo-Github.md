

#How To Github
Okay, that may sound strange, but making a durable pauses I often forget how to do some basic tasks via Git.  
First of all, how to write markdown texts like this one. I guess [Github Tutorial](https://guides.github.com/features/mastering-markdown/) is nice enough, but let's talk commands:


###First Config
```shell
git config --global user.name "Divan"
git config --global user.email "SendMeA@mail.com"

git config --global core.autocrlf input
git config --global core.safecrlf true 
??
git config --global core.quotepath off
```

###Practical stuff
```shell
git init                  # Creates a NEW REPO IN CURRENT FOLDER

git clone *url*           # Downloads from url

git add filename          # Sends file to the query

git commit -m "What A Lovely Day" #Commits, dah?

git push                  # Sends to Hub

git pull                  # Takes updates from server to local

git checkout -b *name*    # Creates a new branch and goes to it

git branch *name*         # Just creates branch

```

###Extra stuff
```shell
git status          # Checks wether there are ambiguos ambivalent stuff

git checkout        # Cancels & drops changes

git reset           # Cancels on 'add' stage

git log --pretty=oneline # dah

git revert HEAD     # Cancels last commit
```
