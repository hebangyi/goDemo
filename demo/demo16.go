package demo

import "fmt"

//// 结构体与方法
// 定义一个Cat结构体
type Cat struct {
	Name  string
	Age   int
	Color string
	hobby string // 小写的变量为私有的
}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              // 指针
	slice  []int             // 切片
	map1   map[string]string // 切片
}

type Monster struct {
	Name string `name` // tag , 在序列化和反序列化作为字段使用
}

//// Go中的结构体
func testStruct() {
	// 当声明一个结构体变量时,结构体的空间就已经分配了
	// 结构体是值类型
	// 创建结构体的方式
	// 1.
	var cat1 Cat
	fmt.Println("cat1=", cat1)
	// 2.
	cat2 := Cat{"小红", 10, "白色", "没有兴趣"}
	fmt.Println("cat2=", cat2)
	// 3 使用new关键字new一个指针
	var cat4 *Cat = new(Cat)
	(*cat4).Name = "smith"
	(*cat4).Age = 30
	fmt.Println(*cat4)
	// 方式4
	var person *Person = &Person{}
	person.Name = "test"

	cat1.Age = 2
	cat1.Name = "小白"
	cat1.Color = "白色"

	fmt.Println("cat1=", cat1)
	fmt.Println("猫的信息如下:")
	fmt.Println("名称=", cat1.Name)
	fmt.Println("年龄=", cat1.Age)
	fmt.Println("颜色=", cat1.Color)

	// 定义结构体变量
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}

	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	// 结构体中的内存地址是连续的
	// 结构体是用户单独定义的类型,和其他类型进行转换时需要有完全相同的字段(名字、个数和类型)

}

type A struct {
}

// 方法绑定
// 其中 a 变量表示当前调用的struct 副本
// 方法的访问控制范围规则和函数一样,方法首字母小写,只能在本包访问,方法首字母大写,可以在本包和其他包访问
func (a A) aMethodTest() {
	// 注意,因为A是struct类型,struct 是值类型,所以该A是值类型的副本
	// a.aMethodTest()
}

// 为了提高效率,我们通常将方法与结构体的指针类型绑定
func (c *A) bMethodTest() {
	// c 是指针,就是该变量
	// 编译器做了优化,等价于用指针的方式进行调用
	c.aMethodTest()
}

// 方法的返回值
func (a A) aReturnTest(param1 int32) int32 {
	return param1
}

// 给*Student实现String方法,打印字符串将调用该方法
func (a *A) String() string {
	return "abc"
}

func testFunc() {
	var a A
	a.aMethodTest()    // 方法调用
	(&a).aMethodTest() // 指针方式调用

	(&a).bMethodTest() // 指针的调用方式
	a.bMethodTest()    // 编译器做了优化,等价于用指针的方式调用
	intVal := a.aReturnTest(1)
	fmt.Println("intVal = ", intVal)

	fmt.Println("A String = ", (&a))
}

/**
Golang中的方法在指定得数据类型上的(即: 和指定的数据类型绑定),因此自定义类型,都可以有方法,而不仅仅是struct, 比如int, float32等都可以有方法
*/
// 修改integer 别名
type integer int

func (i integer) print() {
	fmt.Println("i = ", i)
}

func (i *integer) change() {
	*i = *i + 1
}

func testIntMethod() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println("i = ", i)
}

type Student struct {
	Name string
	Age  int
}

func (stu *Student) printStu() {
	fmt.Printf("Stu Name = %v , Stu Age = %v", stu.Name, stu.Age)
}

func testOop() {
	// 对象实例化的几种方式
	var stu1 Student = Student{"tom", 10}
	fmt.Println("student : ", stu1)

	stu2 := Student{"tom2", 11}
	fmt.Println("student : ", stu2)

	var stu3 Student = Student{
		Name: "mary",
		Age:  30,
	}
	fmt.Println("stu3 = ", stu3)

	stu4 := Student{
		Name: "mary",
		Age:  30,
	}
	fmt.Println("stu4 = ", stu4)
	// 用指针的方式创建
	var stu5 *Student = &Student{"smith", 20}
	fmt.Println("stu5 = ", stu5)

	var stu6 *Student = &Student{Name: "scott",
		Age: 20,
	}
	fmt.Println("stu6 = ", stu6)
}

// 使用工厂方法创建对象
func NewStudent(name string, age int) *Student {
	// 注意,返回的是指针,因为是Struct 类型
	return &Student{
		Name: name,
		Age:  age,
	}
}

func testNewStudent() {
	stu := NewStudent("heby", 27)
	stu.printStu()
}
