package demo

import "testing"

// 单元测试-基本介绍
// go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试
// 可以解决如下问题
// 1.确保每个函数式可运行,并且运行结果是正确的
// 2.确保写出来的代码性能是好的
// 3.单元测试能计算发现程序设计或实现的逻辑错误,使问题及早暴露,便于问题的定位解决,
// 闭包的概念
// go中的testing框架
// 使用 go test -v 命令执行(-v 运行正确或错误都输出日志)
// 将xxx_test.go的文件引入,将Test开始的函数调用
func TestAddUpper(t *testing.T) {
	// 调用
	res := AddUpper()
	if res != nil {
		// 输出错误信息,并退出程序
		t.Fatalf("AddUpper(10) 执行错误,期望值=%v, 实际值=%v\n", 33, res)
	}

	// 如果正确输出日志
	t.Logf("执行正确!")
}
