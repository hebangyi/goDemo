package main

import (
	"fmt"
	"reflect"
)

//// 反射练习

// 反射可以在运行时动态获取变量的各种信息,比如变量的类型(type), 类别(kind)
// 如果是结构体变量,还可以获取到结构体本身的信息(包括结构体的字段,方法)
// 通过反射,可以修改变量的值,可以调用关联的方法
// 需要使用reflect 包

// 对数值类型的反射
func reflectTest01(b interface{}) {
	// 通过反射获取传入的变量的type , kind ,值
	// 1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	// 这里是reflet.Type类型
	fmt.Printf("rVal = %v rVal type = %T", rTyp, rTyp)

	fmt.Println()

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	// 将reflect.Value 导出为Int(该值的原值必须是Int)
	rVal.Int()
	// 返回的Kind是一个常量
	rVal.Kind()

	// 这里是reflect.Value类型
	fmt.Printf("rVal = %v rVal type = %T", rVal, rVal)

	// 将rVal 导出int 转成 interface{}
	iv := rVal.Interface()
	// 将 interface{} 通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2 = ", num2)

}

// 对结构体的反射
func reflectTest02(b interface{}) {
	// 1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	// 这里是reflet.Type类型
	fmt.Printf("rTyp Val = %v rTyp Type = %T \n", rTyp, rTyp)

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	// 这里是reflect.Value类型
	fmt.Printf("rVal Val = %v rVal Type = %T \n", rVal, rVal)

	// 下面我们将 rVal 转成 interface{}
	iv := rVal.Interface()
	// 1. 使用类型断言
	stu, ok := iv.(Student)
	if ok {
		fmt.Println(stu.Name)
	}
	// 2. 使用for 循环
	// for i, x := range iv {
	// 	switch x.(type) {
	// 	case Student:
	// 		fmt.Printf()
	// 	}
	// }

	// 3. 反射结构体中的某个字段
	// 通过某个字段类型获得值
	// 获取结构体中的tag标签
	// rTyp.Field(0).Tag.Get("json")
	// rVal.Field(0).Int()
}

// 通过反射修改值(数值类型)
// Elem 函数
func reflectTest03(b interface{}) {
	// 通过反射获取传入的变量的type , kind ,值
	// 1. 先获取到 reflect.Type
	rVal := reflect.ValueOf(b)
	// 这里是reflet.Type类型
	fmt.Printf("rVal = %v rVal type = %T", rVal, rVal)

	fmt.Println()
	// 这里反射出来的是int指针类型的反射值
	// Elem 返回指针指向的值的反射值
	rVal.Elem().SetInt(101)

	// 反射创建对象
	// reflect.New() // 该值持有一个类型为type的新申请的零值指针
}

// TODO 反射调用方法

type Student struct {
	Name string `json:name`
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	// 反射中变量 interface{} 和 reflect.Value 是可以相互转换的
	fmt.Println("hello")
	// 1.先定义一个int
	var num int = 100
	// reflectTest01(num)

	// 2. 定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  12,
	}
	reflectTest02(stu)

	reflectTest03(&num)
	fmt.Println("reflectTest03 num val = ", num)
}
