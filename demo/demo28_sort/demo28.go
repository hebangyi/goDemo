package demo28

import (
	"fmt"
	"sort"
)

// Pair 代表键值对
type Pair struct {
	Key   string
	Value int
}

func TestSort1() {
	m := map[string]int{
		"apple":  3,
		"banana": 2,
		"orange": 4,
		"grape":  1,
	}

	var pairs []Pair
	// 提取键值对到切片
	for k, v := range m {
		pairs = append(pairs, Pair{k, v})
	}

	// 对切片进行排序
	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	// 打印排序后的map
	for k, v := range pairs {
		fmt.Printf("%v: %v\n", k, m[v.Key])
	}
}

func TestSort2() {
	var intArr []int = []int{0, 3, 2, 1, 4, 8, 6, 7}
	sort.Sort(sort.IntSlice(intArr))
	fmt.Println("intArr = ", intArr)
}

func TestSort3() {
	var intArr []int = []int{0, 3, 2, 1, 4, 8, 6, 7}
	var slice = sort.IntSlice(intArr)
	// 包装反转
	var reverse = sort.Reverse(slice)
	sort.Sort(reverse)
	fmt.Println("reverse = ", reverse)
}

func TestSort4() {
	var intArr []int = []int{0, 3, 2, 1, 4, 8, 6, 7}
	var slice = sort.IntSlice(intArr)
	// 排序
	sort.Sort(slice)
	// 搜索排序后第一个值的索引
	var index = sort.Search(len(slice), func(i int) bool {
		return i >= 6
	})
	fmt.Println("find index = ", index)
}
