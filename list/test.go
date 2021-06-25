package main

import (
	"fmt"
	_ "github.com/intrsokx/leetcode/list/timeutil"
	_ "unsafe"
)

//在编译器中因为没有now()不符合正常语法，所以会飘红报错，但是是不影响执行代码的。
//go语法提供了一种定义方法的机制，即在.go文件中定义方法，在.s汇编代码中进行方法的实现，
//我们可以借助这个特点，在定义方法的包下新建一个空的test.s文件，让编译器误以为该方法是在汇编代码中实现的，从而不飘红。
//go:linkname now github.com/intrsokx/leetcode/list/timeutil.GetNow
func now() string

func main() {
	fmt.Println(now())
}
