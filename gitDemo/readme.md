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
```

