package demo9

/// 流程控制相关
import (
	"fmt"
	"reflect"
)

// TestBranch1 IF 表达式
func TestBranch1() {
	var a int = 10
	var b int = 9
	// IF的条件表达式
	if a > b {
		fmt.Println("a 值 大于 b 值!")
	}
}

// TestBranch2 IF ELSE 表达式
func TestBranch2() {
	if true {

	} else {

	}
}

// TestBranch3 .
func TestBranch3() {
	if true {

	} else if true {

	} else {

	}
}

func TestBranch4() {
	var b = 10
	var a int
	// golang中支持在if中,直接定义一个变量,比如下面的方式
	if a = 11; a > b {
		fmt.Println("a 值 大于 b 值!")
	}
}

// TestSwitch .
func TestSwitch() {
	var param = 1
	// go 中
	// 匹配项后面不需要再加break
	switch param {
	// switch 可以做多条件的
	case 1, 2, 3:
		fmt.Println("1")
	case 4:
		fmt.Println("2")
	case 5:
		{
			fmt.Println("3")
		}
	default:
		fmt.Println("没有匹配到数值字符!")
	}

	switch age := 10; {
	case age == 1:
		fmt.Println("1")
		// switch 穿透
		// fallthrough
	case age == 2:
		fmt.Println("2")
	default:
		fmt.Println("没有找到匹配字符")
	}

	// Type Switch : switch 语句可以被用于type-switch 来判断某个interface变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("X的类型~:%T", i)
	case int:
		fmt.Printf("x 是 int类型")
	case func(int) float64:
		fmt.Printf("x 是 func(int)型")
	case bool, string:
		fmt.Printf("x 是 bool 或者 string型")
	}

	var type1 reflect.Type
	var type2 reflect.Type
	type1 = reflect.TypeOf(x)
	fmt.Println(type1)

	if type1 == type2 {
		fmt.Println(true)
	}
}

// TestFor for循环控制
func TestFor() {
	// for 循环的一种方式
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world!")
	}

	// for 循环的另外一种方式
	var j = 10
	for j <= 11 {
		j++
	}

	// 一种无限循环的方式
	for {
		break
	}

	// 使用 for-range方式遍历
	var str = "abcdef"
	// index是数组返回值的下标
	// value是该下标位置的值
	// 如果不想使用下标index,可以直接把下标index标为下划线_(占位符)
	// index和value的名称不是固定的,可以自行命名
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}
}

// TestGoto Go语言中可以无条件跳转到程序指定行
// 但在设计中,一般不主张使用goto语句,容易造成流程的混乱
func TestGoto() {
	// goto 基本介绍
	fmt.Println("ok1")
	// 跳转到label1
	goto label1
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
label1:
	fmt.Println("ok5")
	fmt.Println("ok6")
}
