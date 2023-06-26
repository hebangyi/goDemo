package demo23

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	fmt.Println(reflect.TypeOf(Student{}).Name())
	fmt.Println("=====", reflect.TypeOf(Student{}).Kind(), "========")
}
