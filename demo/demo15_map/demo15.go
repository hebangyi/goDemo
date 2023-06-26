package demo15

import (
	"fmt"
	"reflect"
)

//// Map 数据结构

// map 创建
// 使用 make 创建
func TestMapCreate() {
	// map声明的举例
	// var a map[string]string
	// var b map[string]int
	// var a map[int]string
	// var d map[string]map[string]string
	// 声明不会分配内存,初始化需要make,分配内存后才能赋值使用
	// 使用make 创建map
	map1 := make(map[string]string)
	map1["name"] = "heby"
	fmt.Println(map1)

	// 10是分配10个map空间
	// map 中的空间如果不够,则会动态扩充map空间
	map2 := make(map[string]string, 10)
	map2["name"] = "test"
	fmt.Println(map2)
}

// 使用自变量进行创建
func TestMapCreate2() {
	map1 := map[string]string{
		"name": "heby",
	}
	fmt.Println(map1)
}

// map 创建和更新
func TestMapAddAndUpdate() {
	map1 := make(map[string]string)
	map1["heby"] = "heby1"
	fmt.Println(map1)
	map1["heby"] = "heby2"
	fmt.Println(map1)
}

// map 删除
func TestMapDelete() {
	map1 := make(map[string]string)
	map1["heby1"] = "heby1"
	map1["heby2"] = "heby2"
	fmt.Println(map1)

	// 删除
	delete(map1, "heby1")
	fmt.Println(map1)
}

// map 查找
func TestFind() {
	map1 := make(map[string]string)
	map1["heby"] = "heby1"

	var val = map1["heby"]
	var val2 = map1["heby2"]
	fmt.Println(val)
	fmt.Println("------------------")
	fmt.Println(reflect.TypeOf(val2))
	fmt.Println(val2 == "")
	fmt.Println("------------------")

	// 还有一种方式 判断Map中存在Key
	var str, findRes = map1["heby3"]
	if findRes {
		fmt.Println("Find Val = ", str)
	} else {
		fmt.Println("不存在 Val")
	}
}

// 多维度map
func TestMaps() {
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["adress"] = "北京长安街"
	fmt.Println(studentMap)
}

// map 遍历
func IteratorMap() {
	var map1 = make(map[string]string)
	map1["heby1"] = "heby1 val"
	map1["heby2"] = "heby2 val"

	for k, v := range map1 {
		fmt.Printf("key = %v , val = %v \n", k, v)
	}
}

func TestMap() {
	var map1 = make(map[string]string)
	TestMap1(map1)
	fmt.Println(map1)
}

func TestMap1(m map[string]string) {
	m["heby3"] = "heby3"
}

func TestSlice() {
	var slice = make([]int, 10)
	TestSlice1(slice)
	fmt.Println(slice)
}

func TestSlice1(slice []int) {
	slice[0] = 1
	slice[1] = 2
}

// map 的长度
func TestMapLen() {
	map1 := make(map[string]string)
	map1["heby"] = "heby"
	fmt.Println("len = ", len(map1))
}

func TestSliceMap() {
	// map 切片的使用
	// 1.声明一个map切片,并初始化
	var monsters []map[string]string
	// 引用空间的使用限制不能超过2
	monsters = make([]map[string]string, 2)
	// 2.增加一个妖怪信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "狐狸精"
		monsters[1]["age"] = "300"
	}

	fmt.Println(monsters)
	newMonster := map[string]string{
		"name": "新的妖怪",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}
