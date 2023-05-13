package demo4

import (
	"fmt"
	"unsafe"
)

// 数值类型
func Test5() {
	// 有符号位的数值类型
	var a int8
	var b int16
	var c int32
	var d int64

	// int 的大小取决于所运行的操作系统和编译器
	fmt.Printf("size of %d ", unsafe.Sizeof(int(10000)))

	fmt.Println(a, b, c, d)
	// 无符号位的数值类型
	var e uint16
	fmt.Println(e)

	// int类型的取值范围
	// 在32位的系统中int类型为4个字节,在64位系统中,为8个字节

	// uint类型的取值范围
	// 在32位的系统中int类型为4个字节,在64位系统中,为8个字节
}

// TestFloat go中的小数类型
func TestFloat() {
	// 单精度的浮点数
	var price float32 = 89.12
	fmt.Println(price)

	// 双精度的浮点数
	var price1 float64 = 89.125
	fmt.Println(price1)

	// Golang 的浮点类型默认是float64 类型
	var num1 = 1.1
	fmt.Printf("num1的数据类型是 %T", num1)

	// 浮点数的表示方式
	// 十进制的形式
	// 科学计数法的形式
	num2 := 5.12
	num3 := .123
	fmt.Println(num2, num3)
	// 5.123 * 10的2次方
	num4 := 5.1234e2
	// 5.123 / 10的2次方
	num5 := 5.1234e-2
	fmt.Println(num4, num5)
}

// TestChar go中字符类型
func TestChar() {
	// char 的存储大小是8个字节
	// 采用asc编码 字符类型 只能包含 ascII型的字符
	//
}

// TestBool go 中的 bool 类型
func TestBool() {
	var b = false
	fmt.Println(b)
}

// TestString go中的string
func TestString() {
	// go的字符串是由单个字节连接起来
	// go中的字符串是使用utf-8标示文本
	// go中的字符串赋值后,字符串修改String 本身的内容

	var str = "hello world!"

	// 可以用``来批量输出字符不需要拼接的字符
	str = `
	"skksdfdsafasdfasdfasd"
	sadfsadfasdf
	`
	fmt.Println(str)
}
