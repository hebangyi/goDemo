package demo17

import "fmt"

// / 继承
type Student1 struct {
	Name  string
	Age   int
	Score int
	class string
}

// 显示他的成绩
func (p *Student1) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", p.Name, p.Age, p.Score)
}

// 继承的方式
type Pupil struct {
	Student1 //这里就是嵌套的匿名结构体Student ,包含Student 中的Name Age Score 属性和方法
	Name     string
}

func PrintStudent(student1 Student1) {
	fmt.Println(student1.Name)
}

func testExtends() {
	pupil := &Pupil{}
	// 调用的时候，层级关系得到了保留
	pupil.Student1.Name = "tom~"
	pupil.Student1.Age = 8
	// 直接调用父类
	pupil.Student1.ShowInfo()
	// 通过子类调用父类
	pupil.ShowInfo()

	// 子类不能转换为父类
	// PrintStudent(*pupil)

	// 复写字段
	fmt.Println(pupil.Name)
	fmt.Println(pupil.Student1.Name)

	// 继承的深入讨论
	// 1.继承的结构体可以使用匿名结构体重的所有的字段和方法，即首字母大写或小写的字段方法都可以使用
	pupil.Student1.class = "1班"
	// 2.匿名结构体字段访问可以简化
	// 编译器先查看子类型是否有该字段，如果有则使用该字段
	// 如果没有就去看B中匿名结构体查找该，如果有就调用，如果没有就继续查找，如果都找不到就报错
	// 3.如果两个匿名结构体有相同的字段和方法，同时结构体本身没有同名的字段和方法，在访问时，就必须明确指定匿名结构体名字，否则编译会报错

	pupil.Name = "smith"
	pupil.Student1.Age = 20
	// 调用继承类的方法
	pupil.Student1.ShowInfo()
	// 4.使用匿名内部类的方式进行赋值
	/*	pupil2 := &Pupil{
			Student1{Name: "张三"},
		}

		fmt.Println(pupil2)*/
}
