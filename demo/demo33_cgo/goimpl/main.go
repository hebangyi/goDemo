package main

//#include <hello.h>
import "C"
import "fmt"

//export SayHello
func SayHello(s *C.char) {
	fmt.Println("this is go method !")
	fmt.Print(C.GoString(s))
}

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
