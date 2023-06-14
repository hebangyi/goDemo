package demo31

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func TestKind(t *testing.T) {
	var count = 10
	for count > 0 {
		count = count - 1
		fmt.Println(count)
	}
}

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func TestPoint(t *testing.T) {
	d1 := data{"one"}
	d1.print() //ok

	// interface 没有实现指针方法
	// var in printer = data{"two"} //error
	// in.print()

	m := map[string]data{"x": data{"three"}}
	// m["x"].print() 这个类型不能直接调用
	d2 := m["x"]
	d2.print() //error

	fmt.Println(m["x"].name)
	// 无法直接调用
	// m["x"].name = "two"
}
