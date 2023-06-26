package demo17

type Parent struct {
	Name string
}

type Child struct {
	*Parent
	Age int
}

func main() {
	/*
		parent := &Parent{Name: "John"}
		child := &Child{
			Parent: parent,
			Age:    10,
		}

		var parentPtr Parent
		// parentPtr = *child          // error 子类型值赋值给父类型变量
		fmt.Println(parentPtr.Name) // 输出：John
	*/
}
