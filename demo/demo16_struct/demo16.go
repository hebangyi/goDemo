package demo16

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

////////////// 创建结构体
// 创建结构体
func TestCreateStruct() {
	// 声明一个struct 结构
	// 零值初始化
	var cat1 Cat
	fmt.Println("cat1=", cat1)
	// 初始化一个struct 结构 1
	cat1 = Cat{"小红", 10, "白色", "吃饭"}
	fmt.Println("cat1=", cat1)
	// 初始化一个struct 结构 2
	cat1 = Cat{Name: "小白", Age: 12}
	fmt.Printf("%v", cat1)
}

// 使用指针的方式创建结构体对象
func TestCreateStructPoint() {

	var cat1 *Cat = new(Cat)
	cat1.Age = 11
	fmt.Println(cat1)
	// 使用&的方式

	cat2 := &Cat{
		Name: "小白",
	}

	fmt.Println(cat2)
}

/// 结构体方法
type A struct {
	Name string
	Age  int
}

// 方法绑定
// 注意: 其中 a 变量表示当前调用的struct 副本
// 方法的访问控制范围规则和函数一样,方法首字母小写,只能在本包访问,方法首字母大写,可以在本包和其他包访问
func (a A) StructModifyName(name string) {
	// 注意,因为A是struct类型,struct 是值类型,所以该A是值类型的副本
	a.Name = name
	a.Age = 20
}

// 为了提高效率,我们通常将方法与结构体的指针类型绑定
func (c *A) PointModifyName(name string) {
	// c 是指针,就是该变量
	// 编译器做了优化,等价于用指针的方式进行调用
	c.Name = name
}

func TestStructModifyName() {
	a := &A{Name: "name1", Age: 30}
	// &{name1 30}
	a.StructModifyName("name2")
	fmt.Println(a)
}

func TestPointModifyName() {
	a := &A{Name: "name1", Age: 30}
	a.PointModifyName("name3")
	fmt.Println(a)
}

// 方法的返回值
func (a A) aReturnTest(param1 int32) int32 {
	return param1
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

func (stu *Student) PrintStu() {
	fmt.Printf("Stu Name = %v , Stu Age = %v", stu.Name, stu.Age)
}

// TOString  方法的复写
func (stu *Student) String() string {
	fmt.Println("这个是复写的String 方法")
	return "print student"
}

// 使用工厂方法创建对象
func NewStudent(name string, age int) *Student {
	// 注意,返回的是指针,因为是Struct 类型
	return &Student{
		Name: name,
		Age:  age,
	}
}

func TestNewStudent() {
	stu := NewStudent("heby", 27)
	stu.PrintStu()
	fmt.Println(stu)
}
