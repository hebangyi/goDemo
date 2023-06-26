package demo4

import (
	"fmt"
	"reflect"
	"testing"
)

// / 接口
// / 声明一个接口
type Usb interface {
	//声明两个没哟实现的方法
	Start()
	Stop()
}

type Start interface {
	Start()
}

type Stop interface {
	Stop()
}

type Phone struct {
}

// 让Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 让Camera 实现 Usb 接口的方法
type Camera struct {
}

// 让Camera 实现Usb 接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

// 计算机
type Computer struct {
}

// 编写一个方法Working 方法,接受一个Usb接口类型变量
// 只要是实现了 Usb 接口(所谓实现Usb接口,就是指实现了 Usb 接口声明所有方法)
func (c Computer) Working(usb Usb) {
	// 通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

type Object interface {
	Say()
}

type integer1 int

func (i integer1) Say() {
	fmt.Println(i)
}

type interfaceA interface {
	MethodA()
}

// 接口继承
type interfaceB interface {
	interfaceA
	MethodB()
}

// 泛型约束 (这种叫一般接口 前者是基本接口)
type InterfaceC interface {
	// ~ 指定底层类型
	float32 | float64 | ~int
}

type Person[T InterfaceC] struct {
}

// 空接口,可以接受任意类型
type empty interface {
}

// 断言判断
func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) { // 这里的type是一个固定的关键字,返回类型
		case bool:
			fmt.Printf("param #%d is a bool 值是%v", i, x)
		case int:
			fmt.Printf("param #%d is a int 值是%v", i, x)
		}
	}
}

type Animal interface {
	Sleep()
	// Talk()
}

type Student struct {
	Name string
}

func (s *Student) Sleep() {
	fmt.Println(s.Name + "学生在睡觉")
	s.Name = "fin heby"
}

func OnSleep(animal Animal) {
	animal.Sleep()
}

func TestInterface(t *testing.T) {
	var use Usb = Phone{}
	fmt.Println(reflect.TypeOf(use))
	_, ok := use.(Start)
	fmt.Println(ok)
}

func TestExtends(t *testing.T) {
	stu := Student{Name: "heby"}
	OnSleep(&stu)
	fmt.Println(stu.Name)
}

func TestMethod(t *testing.T) {
	//测试
	// 先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	// 接口调用
	// Usb 中的接口必须全部实现
	computer.Working(phone)
	computer.Working(camera)

	// 接口不能创建实例,但是可以指向一个实现该接口的变量
	var u Usb = Phone{}
	fmt.Println(u)

	// 不一定是结构体,非结构体实现了接口方法也可以指向
	var i integer1 = 10
	var obj Object = i
	obj.Say()

	// 类型断言
	var a interface{}
	var computerA Computer = Computer{}
	a = computerA
	var computerB Computer
	// 类型断言
	// 表示判断a是否指向Point类型的变量,如果是就转成Point类型并赋值给b变量,否则报错
	computerB = a.(Computer)
	fmt.Println(computerB)

	// 有错误处理的断言语法工规则
	if computerC, ok := a.(Computer); ok {
		fmt.Println("convert success")
		fmt.Println(computerC)
	} else {
		fmt.Println("convert fail")
	}
}
