package demo

import (
	"fmt"
	"unsafe"
)

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
