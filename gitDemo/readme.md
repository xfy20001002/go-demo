# git 
## 初始化
git init
git add filename
git commit -m "commit information"
## 查看仓库状态
git status  :显示哪个文件被修改,但还没有被提交
git diff filename :查看修改的文件名
## 版本回退
git log :查看提交的版本
git log --pretty=oneline :查看提交版本，并在一行中显示
```
用HEAD表示当前版本，上一个版本就是HEAD^，上上一个版本就是HEAD^^，当然往上100个版本写100个^比较容易数不过来，所以写成HEAD~100

命令
回退到上一个版本
git reset --hard HEAD^
从上一个版本回到现在版本
git reset --hard [现在版本的commitID] //从命令行中去找
git reflog //查看命令历史(可以找到各个命令commit版本ID)
```

## 撤销修改
git checkout -- <filename> //撤销该文件的修改(非暂存区)
git reset HEAD <file> //把暂存区的修改撤销

## 删除文件
场景1
工作区删除了文件 版本库同样要删除
```
git rm <filename>
git commit -m "remove test.txt"

```
场景2
删错了文件，从版本库中恢复
```
git checkout -- <filename>
```
git checkout其实是用版本库里的版本替换工作区的版本，无论工作区是修改还是删除，都可以“一键还原”

## 远程仓库
场景1
本地建立git仓库后想要与远程仓库关联
```
1.远程建立一个空的仓库
2.本地输入如下命令(本地与远程库相关联)
git remote add origin git@github.com:michaelliao/learngit.git
3.将本地库中所有内容推送到远程库
git push -u origin master

加上-u参数，Git不但会把本地的master分支内容推送的远程新的master分支，还会把本地的master分支和远程的master分支关联起来
```
场景2
删除远程库
```
git remote -v 查看远程库信息
git remote rm origin //删除名为origin的远程库
```

## 创建并切换分支
```
git checkout -b dev = git branch dev + git checkout dev
-b表示创建并切换

git merge <branchname> 
将目标分支(branchname)工作成果合并到master上

删除分支
git branch -d <branchname>
git branch -D <name> 丢弃一个没用合并过的分支

合并分支(两个分支在同一文件发生修改之后合并会发生冲突的情况 需要手动修改再提交)

不用fast forward形式合并分支
git merge --no-ff -m "merge with no-ff" dev

修复bug
1.git stash将工作现场隐藏起来
2.确定要在哪个分支上修复bug，假定需要在master分支上修复，就从master创建临时分支
3.恢复工作现场 
git stash list 查看stash区内容
git stash apply(恢复工作现场 stash中内容不删除)
git stash pop(恢复工作现场 stash中内容删除)
git stash apply stash@{0} 在有多个stash时恢复指定工作现场

复制某个特定的提交到当前分支
git cherry-pick <commitId>
```

## 远程分支
```
远程库默认名称为origin
git remote 查看远程库名称 -v参数显示更详细的信息

推送分支
git push origin <branchname> 把该分支上的所有本地提交推送到远程库
```

## 抓取分支
```
git clone git@github.com:michaelliao/learngit.git
创建远程的dev分支到本地
git checkout -b dev origin/dev

推送分支
git push origin dev 若失败表示远程分支有更新
git pull 若失败表示没有设置链接关系
git branch --set-upstream-to=origin/dev dev 设置链接关系
若有冲突手动合并冲突

```

