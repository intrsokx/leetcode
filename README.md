# leetcode

### reference
* [tree](./tree/tree.md)
* [list](./list/list.md)
* [binary](./binary/binary.md)


### 获取项目git信息，在项目编译阶段注入到程序中
```shell script
commitId=`git rev-parse --short HEAD`
name=`git log --pretty=format:“%an” $commitId -1`
date=`git log --pretty=format:“%cd” $commitId -1`
message=`git log --pretty=format:“%s” $commitId -1`

go build -ldflags "-X main.VERSION=$commitId -X main.NAME=$name -X 'main.DATE=$date' -X 'main.MESSAGE=$message'" test.go

#"程序正常执行"
./test
ehco  ""
#"打印当前程序（二进制可执行文件）在github上提交的信息"
./test --info=true
```