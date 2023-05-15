// 如果要编译生成一个可执行程序文件,就需要将这个包声明为main 即 package main
package main

import demo20 "example.com/go-demo-1/demo/demo20_json"

func main() {

	// name := demo.GetName()
	// fmt.Println(name)
	// fmt.Println(demo.GetName())
	// fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	// demo.Test()
	// demo.Basictype2ConvertString()
	//fmt.Println("hello world");
	// demo.Basictype2ConvertString()

	// fmt.Println("hello world" + " and go!")
	// demo3.Test3()

	// 如果需要修改字符串,可以先将
	// demo4.Test5()

	// num := 123
	// fmt.Println(reflect.TypeOf(num))
	// demo13.TestError()
	// demo13.TestPanic()

	// demo14.TestSliceFunc()
	// demo15.TestMap1()
	// demo15.TestMapAddAndUpdate()
	// demo15.TestFind()
	// fmt.Println("这个是 main 函数")
	// demo15.IteratorMap()
	// demo15.TestMapLen()
	// demo14.TestSlice2()
	// demo15.TestSliceMap()
	// demo16.TestCreateStruct()
	// demo16.TestPointModifyName()
	// demo16.TestNewStudent()
	// demo20.TestStructToJson()
	// demo14.TestArray()
	// demo28.TestSort2()
	// demo28.TestSort3()
	// demo28.TestSort4()
	// demo23.TestReflectTypeAndVal(Person{Name: "heby"})
	// demo23.TestReflectTypeAndVal(1)
	// demo23.TestStruct()
	// demo23.TestStruct()
	demo20.TestStructToJson()
}
