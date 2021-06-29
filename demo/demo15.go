package demo

import "fmt"

//// Map 数据结构
//
func testMap() {
	// map声明的举例
	// var a map[string]string
	// var b map[string]int
	// var a map[int]string
	// var d map[string]map[string]string
	// 声明不会分配内存,初始化需要make,分配内存后才能赋值使用

	// map的声明和注意的事项
	// 1.先声明,再make
	var aMap map[string]string
	// 10是分配10个map空间
	// map 中的空间如果不够,则会动态扩充map空间
	aMap = make(map[string]string, 10)
	// 2.声明的过程中直接make
	b := make(map[string]string)

	// map的增加与更新
	b["cc"] = "cc"
	// map的删除
	delete(b, "cc")
	// map的查找
	var val, findRes = b["cc"]
	if findRes {
		fmt.Println("找到了,value=", val)
	} else {
		fmt.Println("value没有找到!")
	}

	// 3.声明,直接赋值,这种方式并没有直接make , 但是底层使用了make方式进行初始化
	var c map[string]string = map[string]string{
		"no4": "成都",
	}
	fmt.Println("c=", c)

	aMap["no1"] = "AAA"
	// map 中的key不能重复
	aMap["no1"] = "BBB"
	fmt.Println("map value = ", aMap)

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["adress"] = "北京长安街"

	fmt.Println(studentMap)

	// map的遍历只能使用for-range变量
	for mapKey, mapVal := range aMap {
		// 分别是 key value
		fmt.Printf("map key = %v , map value = %v \n", mapKey, mapVal)
	}

	// 统计map的key-value 长度
	fmt.Println("len map = ", len(studentMap))

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
}
