package demo23

import (
	"fmt"
	"reflect"
)

//// 反射练习

// 反射可以在运行时动态获取变量的各种信息,比如变量的类型(type), 类别(kind)
// 如果是结构体变量,还可以获取到结构体本身的信息(包括结构体的字段,方法)
// 通过反射,可以修改变量的值,可以调用关联的方法
// 需要使用reflect 包
func TestReflectTypeAndVal(b interface{}) {
	// 先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	// 这里是reflet.Type类型
	fmt.Printf("rType val = %v rType type = %T \n", rType, rType)

	// 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	// 这里是reflect.Value类型
	fmt.Printf("rVal val = %v rVal type = %T \n", rVal, rVal)

	// 获取 Val的类型
	fmt.Printf("rVal kind = %v \n", rVal.Kind())

	// 将val转换为interface
	iv := rVal.Interface()
	rVal.Int()
	num, ok := iv.(int)
	if ok {
		fmt.Println("this is int = ", num)
	} else {
		fmt.Println("this is not int = ", num)
	}
}

func TestStruct() {
	testStructNew()
	testStructTag()
}

func testStructNew() {
	// 使用反射创建一个对象
	// New函数创建的是指向实例类型的指针
	p := reflect.New(reflect.TypeOf(Student{}))
	// 实例化类型对象指针
	var stu = p.Interface().(*Student)
	stu.Name = "heby"
	fmt.Println(stu)

	// 两种方式
	switch p.Interface().(type) {
	case *Student:
		fmt.Printf("this is student")
	}

}

func testStructTag() {
	rType := reflect.TypeOf(Student{})
	for i := 0; i < rType.NumField(); i++ {
		var field = rType.Field(i)
		fmt.Println("field name = ", field.Name, "field type = ", field.Type, "field type = ", field.Tag)
	}

	// 查找固定的标签
	for i := 0; i < rType.NumField(); i++ {
		var field = rType.Field(i)
		fmt.Println("json tag = ", field.Tag.Get("json"))
	}
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

// TODO Elem

type Student struct {
	Name string `json:"name" xml:"xxx"`
	Age  int    `json:"1"`
}

type Monster struct {
	Name string
	Age  int
}
