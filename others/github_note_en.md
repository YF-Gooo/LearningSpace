## git clone
    git clone

## create a new git project
    mkdir learn
    cd learn
    git init
    touch 1.txt
    git status
    git add .
    git commit -m "first commit"


## create branch
    git branch -a //check all branches in local
    git checkout dev //create dev branch
    git checkout master //switch back to master branch
    git merge dev //merge dev to master

## upload to github 
    git remote add origin https://github.com/YF-Gooo/LearningSpace.git
    git push --set-upstrean origin master
    git checkout dev //upload dev branch
    git push --set-upstream origin dev

## git config
    //set global name
    $ git config --global user.name “github’s Name”
    $ git config --global user.email "github@xx.com"
    $ git config --list

    //set local name
    $ git config user.name “gitlab’s Name”
    $ git config user.email "gitlab@xx.com"
    $ git config --list


