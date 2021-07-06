commitId=`git rev-parse --short HEAD`
name=`git log --pretty=format:“%an” $commitId -1`
date=`git log --pretty=format:“%cd” $commitId -1`
message=`git log --pretty=format:“%s” $commitId -1`


#替换message中的'，因为message中包含'会导致go build -ldflags ”-X 'main.MESSAGE=$message'“失败
#因为在-ldflags构建参数中'有特殊含义，所以message中不能包含'
old="'"
new=""
message=${message//$old/$new}

go build -ldflags "-X main.VERSION=$commitId -X main.NAME=$name -X 'main.DATE=$date' -X 'main.MESSAGE=$message'" test.go

#"程序正常执行"
./test --msg='hello world'
echo ""
#"打印当前程序（二进制可执行文件）在github上提交的信息"
./test --info=true
