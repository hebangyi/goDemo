package main

/*
#include <stdio.h>
#include <hello.h>

static void SayHello(const char* s) {
    puts(s);
}

*/
import "C"

func main() {
	Test3SayHello()
}

func Test1HelloWorld() {
	// 使用 c 中 stdio.h 标准库
	C.puts(C.CString("Hello, World\n"))
}

func Test2SayHello() {
	// 使用自定义函数
	C.SayHello(C.CString("Hello, World\n"))
}

func Test3SayHello() {
	// 模块化引入
	C.SayHello1(C.CString("Hello, World\n"))
}
