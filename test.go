package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	VERSION, NAME, DATE, MESSAGE string
	//git 字符串的标志
	pre  = `“`
	suff = `”`
)

var isInfo = flag.Bool("info", false, "打印程序版本信息")

func main() {
	flag.Parse()
	if *isInfo {
		gitDateFmt := pre + `Mon Jan 02 15:04:05 2006 -0700` + suff
		dateFmt := pre + `2006-01-02 15:04:05` + suff
		date, err := time.Parse(gitDateFmt, DATE)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("version:%s\tname:%s\tdate:%s\tmessage:%s", VERSION, NAME, date.Format(dateFmt), MESSAGE))
		return
	}

	fmt.Println("执行正常程序流程...")
}

//test.sh
/**
commitId=`git rev-parse --short HEAD`
name=`git log --pretty=format:“%an” $commitId -1`
date=`git log --pretty=format:“%cd” $commitId -1`
message=`git log --pretty=format:“%s” $commitId -1`
go build -ldflags "-X main.VERSION=$commitId -X main.NAME=$name -X 'main.DATE=$date' -X 'main.MESSAGE=$message'" test.go
*/
