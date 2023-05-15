package demo3

import "fmt"

// 运算 & 拼接
// Test1 + 号做加法运算
func Test1() {
	var i = 1
	var j = 2
	var r = i + j
	fmt.Println("r = ", r)
}

// Test2 + 号对字符串的拼接
// 如果是字符串类型可以使用 + 拼接
func Test2() {
	var str1 = "hello"
	var str2 = "world"
	var res = str1 + str2
	fmt.Println("res = ", res)
}

func Test3() {
	var num1 = 10
	var str = "hello world!"
	var newStr = fmt.Sprintf("%d "+str, num1)
	fmt.Println(newStr)
}
