package demo5

import (
	"fmt"
	"strconv"
	"unsafe"
)

// 类型转换
// Test6 查看变量的数据类型和变量的数据类型大小
func Test6() {
	// 有符号位的数值类型
	var i int = 100
	// 打印数据类型
	fmt.Printf("i 的类型是 %T\n", i)
	// 打印数据占用的字节大小
	fmt.Printf("n2 占用的字节数是 %d", unsafe.Sizeof(i))
}

// TestConvert golang 转换
func TestConvert() {
	// GO转换时需要显示转换,不能自动转换

	var i int32 = 100
	// 将i转为float类型
	var j float32 = float32(i)

	fmt.Println(j)
}

// 格式化函数
// Basictype2ConvertString 基本数据类型转字符串
func Basictype2ConvertString() {
	var num1 int = 99
	// var num = 23.456
	// var b = true
	// var myChar byte = 'h'
	var str string //空的str

	// 使用 fmt.Sprintf 方法来进行转换
	str = fmt.Sprintf("%d haha", num1)
	fmt.Printf("str type %T str = %v", str, str)

	// 使用 strconv包的函数 第二个参数是表示10进制
	strconv.FormatInt(int64(num1), 10)

	// 使用 Itoa 等效于 FormatInt(int64(num1), 10)
	strconv.Itoa(num1)
}

// String2ConvertBasictype string类型转换为基本数据类型
func String2ConvertBasictype() {
	var str = "123456"
	var n1 int64
	// 第二个参数表示10进制 第三个参数表示64位
	n1, _ = strconv.ParseInt(str, 10, 64)
	fmt.Printf("n1 type %T n1=%v", n1, n1)

	// strconv.ParseFloat()
	// strconv.ParseBool()
	// strconv.ParseFloat()
}
