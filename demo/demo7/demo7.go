package demo7

// 面向对象相关
// 函数作用域
var age int = 50

// Name 变量表示变量能否直接访问
var Name string = "jack"

// NewName := "new jack" // 不能使用:=在函数外赋值变量

//TestPublicMethod go中通过首字母是否大写来表示方法是否能直接访问

func TestPublicMethod() {

}

// 小写表示私有方法
func testPrivateMethod() {

}
