package demo26

import "fmt"

// 在引用 demo26 package 的时候 会自动调用init 函数
func init() {
	fmt.Println("init method...")
}

func Run() {
	fmt.Println("run ...")
}
