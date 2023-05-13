package demo2

import "fmt"

// 常量&变量与方法
// Test1 变量声明
func Test1() {
	// 变量的声明本质上已经赋值该变量的内存大小

	// 第一种:指定变量类型
	var i int
	i = 10
	fmt.Println("i = ", i)

	// 第二种:使用类型推导的方式
	var num = 10.11
	fmt.Println("num=", num)

	// 第三种:省略var , 注意 :=左侧的变量不应该是已经声明过的,否则会导致编译错误
	// := 等价与声明变量
	name := "tom"
	fmt.Println("name=", name)

	// 数值的声明方式
	var byteArr []byte
	byteArr = []byte{97, 98}
	fmt.Printf("byteArr = %v", byteArr)
}

// Test2 多变量声明
func Test2() {
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 一次性声明多个变量方式
	var x1, x2, x3 = 100, 1001, 1002
	fmt.Println("x1=", x1, "x2=", x2, "x3=", x3)

	// 渠道var变量
	y1, y2, y3 := 100, 1001, 1002
	fmt.Println("y1=", y1, "y2=", y2, "y3=", y3)
}

//TestInput 输入测试
func TestInput() {
	// 测试输入

	var name string
	// 读取语句,赋值到文件地址
	fmt.Scanln(&name)
	//
	fmt.Scanf("%s", &name)
}

//TestNum 测试二进制 八进制 十六进制
func TestNum() {
	//八进制 以数字0开头表示
	var j int = 011 // 011=>9
	fmt.Println("j=", j)

	// 十六进制 以0x或0X开头表示
	var k int = 0x11 // 0x11=>
	fmt.Println("k=", k)
}

// 多参数声明
func TestMethod(num ...int) {

}

func testConst() {
	// 常量声明的时候必须给值
	// 常量不能修改
	// 常量只能修饰 bool 数值类型 string类型
	const tax int = 0
}
