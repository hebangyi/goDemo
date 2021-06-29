package demo

import (
	"encoding/json"
	"fmt"
)

// Json 序列化与反序列化
// 定义一个结构体

type JsonMonster struct {
	Name     string `json="monster_name"` // 使用别名
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStructToJson() {
	// 演示
	monster := JsonMonster{
		Name: "牛魔王",
		Age:  10,
	}

	// 将monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 error=%v", err)
	}

	// 输出序列化结果
	fmt.Printf("monster 序列化后=%v", string(data))
	// 反序列化
	// 定义一个Monster实例
	var remonster JsonMonster
	json.Unmarshal(data, &monster)
	if err == nil {
		fmt.Println("monster = ", remonster)
	}

}

func testMapToJson() {
	// 定义一个map
	var a map[string]interface{}
	// 使用map,需要make
	a = make(map[string]interface{})
	a["name"] = "honghaiwer"
	a["age"] = 13
	a["address"] = "洪崖洞"

	// 将a这个map进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 error=%v", err)
	}

	// 输出序列化结果
	fmt.Printf("monster 序列化后=%v", string(data))

	// 反序列化
	// 注意: 反序列化map,不需要make,因为make操作被封装到Unmarshal函数中
	err1 := json.Unmarshal(data, &a)
	if err1 != nil {
		fmt.Printf("unmarsha1 err = %v\n", err)
	}
}

func testSliceJson() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	// 使用map前需要make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = "重庆"
	slice = append(slice, m2)

	// 将切片进行序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 error=%v", err)
	}

	// 输出序列化结果
	fmt.Printf("monster 序列化后=%v", string(data))
}
