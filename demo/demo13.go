// golang中的错误处理机制,异常处理机制

package demo

import (
	"errors"
	"fmt"
)

// go中的异常处理
// go中处理异常的方法是defer,panic,recover
func testError() {
	// 使用defer + recover 来处理错误
	defer func() {
		err := recover() // recover()内置函数,可以捕获到异常
		if err != nil {  // 如果捕获到错误
			fmt.Println("err=", err)
		}
	}() // 调用匿名函数

	num1 := 10
	num2 := 0
	res := num1 / num2

	fmt.Println("res=", res)
}

// 自定义错误
func testErrorNew(fileName string) (err error) {
	if fileName == "testFile" {
		return nil
	}
	// 返回一个自定义错误
	return errors.New("读取文件错误")

}

// 测试自定义错误的使用
func normalDefine() {
	err := testErrorNew("testFile1")
	if err != nil {
		// 如果程序发送错误,就输出错误,并终止程序
		// panic 可以接受一个interface{}类型的值, 也可以接受error 类型的变量
		panic(err)
	}

	fmt.Println("后面代码执行......")
}
