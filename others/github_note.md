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
    git checkout -b 
    dev //创建dev分支
    git checkout master //切换master分支
    git merge dev //把dev分支合并进入(master)

## 删除分支
    git branch -d 分支名
    git remote prune origin //如果github上面的分支已经删掉了，需要运行这个删除本地的跟踪

## github 上传
    //把项目上传到github
    git remote add origin https://github.com/YF-Gooo/LearningSpace.git
    git push --set-upstrean origin master
    git checkout dev //上传dev分支
    git push --set-upstream origin dev

## git config
    Git全局配置和单个仓库的用户名邮箱配置

    学习git的时候, 大家刚开始使用之前都配置了一个全局的用户名和邮箱

    $ git config --global user.name “github’s Name”
    $ git config --global user.email "github@xx.com"
    $ git config --list

    如果你公司的项目是放在自建的gitlab上面, 如果你不进行配置用户名和邮箱的话, 则会使用全局的, 这个时候是错误的, 正确的做法是针对公司的项目, 在项目根目录下进行单独配置

    $ git config user.name “gitlab’s Name”
    $ git config user.email "gitlab@xx.com"
    $ git config --list

