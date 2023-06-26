package demo16

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestStruct(t *testing.T) {
	var cat Cat
	cat.Name = "heby"
	cat.Color = "orange"
	fmt.Println(cat)

	var arr [3]int
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	fmt.Println(arr)

	fmt.Println(unsafe.Sizeof(arr))
	fmt.Println(unsafe.Sizeof(Cat{}))

	// typ := reflect.TypeOf(Cat{})
	// ins := reflect.New(typ).Elem()
}
