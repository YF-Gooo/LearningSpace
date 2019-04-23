## git 创建新项目
    mkdir learn
    cd learn
    git init
    touch 1.txt
    git status
    git add .
    git commit -m "first commit"
    
## git branch 创建分支
    git branch -a //查看所有分支
    git checkout dev //创建dev分支
    git checkout master //切换master分支
    git merge dev //把dev分支合并进入(master)

## github 上传
    //把项目上传到github
    git remote add origin https://github.com/YF-Gooo/LearningSpace.git
    git push --set-upstrean origin master
    git checkout dev //上传dev分支
    git push --set-upstream origin dev