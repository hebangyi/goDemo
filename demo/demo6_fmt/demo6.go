package demo6

import "fmt"

// 指针测试
func TestPoint() {
	// 基本数据类型在内存中布局
	var i int = 10
	// i的内存地址是什么 , &i
	fmt.Println("i的地址=", &i)

	// 下面的 var ptr *int = &i
	// 1. ptr是一个指针变量
	// 2. ptr的类型 *int
	var ptr *int = &i
	fmt.Printf("ptr 中存取的内存地址=%v\n", ptr)
	fmt.Printf("&ptr 存放指针空间的内存地址=%v\n", &ptr)
	fmt.Printf("*ptr 中地址指针指向的值=%v\n", *ptr)
	// 3.对指针变量中对应的地址赋值,指针赋值的类型必须是对应的类型
	*ptr = 211
	// 4. 值类型都有对应的指针类型
	fmt.Println(fmt.Sprintf("Repeated Set Parent : #{123}"))
}
