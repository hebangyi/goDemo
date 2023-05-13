package demo10

// 包、函数相关
import (
	// 使用包的别名
	"fmt"
	f "fmt"
)

var (
	// Fun1 就是一个全局匿名函数，大写表示可以全局调用，和方法一样
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

// 在引用 demo10 package 的时候 会自动调用init 函数
// init 函数中可以完成函数文件的初始化工作
func init() {
	fmt.Println("init()...")
}

// TestFunc1 一般函数
func TestFunc1(a int32) {
}

// TestFunc2 函数的形参传入与返回
func TestFunc2(a int32) (int32, int32) {
	return a, 2 * a
}

// TestFunc3 在Go中,函数也是一种数据类型
// 可以将函数赋值给一个变量
func TestFunc3() {
	var a = TestFunc1
	fmt.Printf("%T", a)
	// 调用函数
	a(123)
}

// testFunc4在go中定义类型别名
func testFunc4() {
	// 给int 类型取别名 myInt 代表 int 类型
	type myInt int
	var a myInt = 123
	fmt.Println(a)
	// 如果int类型
	var b int = int(a)
	fmt.Println(b)
}

// 函数作为形参进行调用
func myFun(function func(int32) (int32, int32), num1 int32, num2 int32) (int32, int32) {
	return function(num1)
}

// 给返回值赋值
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 支持0到多个参数
// 其中args是slice切片,通过args[index]可以访问各个值
func sum(args ...int) (sum int) {
	return 0
}

// 支持1到多个参数
func sum2(n1 int, args ...int) (sum int) {
	return 0
}

// TestFunc 测试函数调用
func TestFunc() {
	TestFunc1(123)

	// 希望忽略某个返回值,使用_符号表示占位忽略
	var res, _ = TestFunc2(123)
	f.Println(res)
}

// TestFuncAnony 匿名函数
func TestFuncAnony() {
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	res2 := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Println(res1)
	fmt.Println(res2(20, 10))
}

// AddUpper 匿名函数的闭包概念测试
func AddUpper() func(int) int {
	var n int = 10
	// 匿名函数和外部的变量形成整体形成闭包
	return func(x int) int {
		n += x
		return n
	}
}

// TestDefer 函数中的defer
// 在函数执行完成后快速释放资源
func TestDefer(n1 int, n2 int) {
	// 1.当执行到defer时,暂时不执行,会将defer后面的语句压入到独立的栈(defer栈)
	// 2.当函数执行完毕后,再从defer栈,按照先入后出的方式出栈,执行
	// 3.在defer将语句放入到栈中时,将相关的值拷贝同时入栈
	defer fmt.Println("defer function 1!")
	defer fmt.Println("defer function 2!")
	defer func() {
		fmt.Println("defer 函数调用")
	}() // 调用匿名函数
	fmt.Println("normal function!")
}

// 内建函数
func testInnerFunc() {
	// len的使用 数组的长度
	// new的使用
	// num2的值是指针类型的值
	num2 := new(int) // *int
	fmt.Printf("num1的类型%T , num1的值 = %v ,num1的地址 %v \n", num2, num2, &num2)

	// make
}
