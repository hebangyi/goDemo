// golang中的错误处理机制,异常处理机制

package demo13

import (
	"errors"
	"fmt"
)

// go中的异常处理
// go中处理异常的方法是defer,panic,recover
func TestError() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常捕获!")
		}
	}()

	// 使用defer + recover 来处理错误
	num1 := 10
	num2 := 0
	res := num1 / num2

	fmt.Println("res=", res)
	fmt.Println("后面的代码继续执行...")
}

// 自定义错误
func New(fileName string) (err error) {
	if fileName == "testFile" {
		return nil
	}
	// 返回一个自定义错误
	return errors.New("读取文件错误")

}

// 测试自定义错误的使用
func NormalDefine() {
	err := New("testFile1")
	// 如果panic 没有通过recover捕获 会终止程序
	// panic 可以接受一个interface{}类型的值, 也可以接受error 类型的变量
	// 请注意，defer 和 recover 是 Go 中用于处理 panic 和异常情况的机制，但在通常情况下，Go 鼓励使用返回错误值来处理错误和异常情况，而不是滥用 panic 和 recover。
	// 只有在非常特殊的情况下，才应该使用 panic 和 recover。
	panic(err)
}

func TestPanic() {
	defer func() {
		fmt.Println("进入 TestPanic")
		if err := recover(); err != nil {
			fmt.Printf("%v \n", err)
		}
	}()
	NormalDefine()
	fmt.Println("后面代码执行......")
}
