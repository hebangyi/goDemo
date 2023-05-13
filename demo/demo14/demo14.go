// go语言的数组与切片

package demo14

import "fmt"

// 数组是多个相同类型数据的组合,一个数组一旦声明/定义,其长度是固定的,不能动态变化
// 数组中的数据类型不能混用
// 数组创建后,如果没有赋值,有默认值
// 使用数组的步骤 1.声明数组并开辟空间 2.给数组各个元素赋值 3.使用数组
// go中的数组属于值类型,默认情况下是值穿点,因此会进行 值拷贝 ,数组间不会相互影响
// 如想在其它函数中,去修改原来的数组,可以使用引用传递(指针的方式)
// 长度是数组类型的一部分
func testArray() {
	// 1.定义一个长度为6的数组
	var hens [6]float64
	// 2.给数组中的元素赋值
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	// 数组的遍历
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}

	// index 是数组的下标
	for index, value := range hens {
		fmt.Println("index = ", index, " value = ", value)
	}

	weight := fmt.Sprintf("%.2f", totalWeight)
	fmt.Println("weight = ", weight)

	// 数组初始化方式
	var numArray0 [3]int
	var numArray1 [3]int = [3]int{1, 2, 3}
	var numArray2 = [3]int{8, 9, 70}
	numArray3 := [3]int{1, 2, 3}
	// 这里的[...]是规定的写法,长度由初始化的值规定
	var numArray4 = [...]int{8, 9, 10}
	// 按照下标的方式进行初始化,长度由初始化的值规定
	var numArray5 = [...]int{1: 800, 0: 900, 2: 999}

	fmt.Printf("%v,%v,%v,%v,%v,%v", numArray0, numArray1, numArray2, numArray3, numArray4, numArray5)
}

// 切片相关
// 切片是数组的一个引用,因此切片是引用类型,在进行传递时,遵守引用传递的机制
// 切片的使用和数组类似,遍历切片、访问切片元素和求切片长度都一样
// 切片的长度是可以变化的,因此切片是一个动态变化的数组

////// 演示切片的基本使用

////// 在内存上讲,slice由3个部分构成,切片的首地址,切片的长度,切片的空间
// 从底层还是说,其实slice就是一个数据结构
// type slice struct{
// 	ptr *[2] int
// 	len int
// 	cap int
// }
func TestSlice() {
	////// 创建切片的几种方式
	//// 声明/定义一个切片
	//// 第一种
	// slice := intArr[1:3]
	// 1. slice 就是切片名
	// 2. intArr[1:3] 表示 slice 引用到intArr这个数组
	// 3. 引用intArr数组的起始下标为1 , 最后的下标为3(但是不包含3)
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 数组和切片的关系是,数组声明对应的是初始化,切片是对数组进行引用
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Printf("slice=%v", slice)
	fmt.Println("slice 的元素是 =", slice)
	fmt.Println("slice 的元素个数是 =", len(slice))
	// cap是一个内置函数,用于统计切片得容量,即最大可以存放多少个元素
	fmt.Println("slice 的容量 =", cap(slice))

	// 如果越界方位会触发异常
	// fmt.Println(intArr[1000])
}

func TestSlice2() {
	// 需要注意的是，如果对切片进行追加操作，可能会触发重新分配底层数组的内存空间，导致切片指向新的底层数组，而不再与原始数组共享。
	// 在这种情况下，原始数组和其他切片的值将不受影响。(其他切片还是引用的原数组)

	// 声明 长度和空间都为10
	var slice0 []int = make([]int, 10)
	fmt.Println("slice0 size = ", len(slice0))
	fmt.Println("slice0 cap =", cap(slice0))

	//// 第二种
	// 通过make方式创建切片
	// make默认创建了一个数组,程序员是不可见的,切片在底层进行维护
	// 使用切片类型,所使用的声明是[],不规定数量
	var slice1 []int = make([]int, 4, 10) // 使用make换算创建切片,第一个参数是类型,第二个参数是大小(Len),第三个参数是切片的空间长度()

	// 给切片得各个元素赋值
	slice1[1] = 10
	slice1[3] = 20
	fmt.Println(slice1)
	fmt.Println("slice1的size=", len(slice1))
	fmt.Println("slice1的cap=", cap(slice1))
	// 只能访问Len所在的空间
	// fmt.Println(slice1[6])

}

func TestSlice3() {
	//// 第三种
	// 直接使用数组的方式
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("cap = ", cap(intArr))
	var slice2 = intArr

	///////////////////////////////////////////////////////////////////////

	///////////////////////////////////////////////////////////////////////

	// 前片的内置函数
	fmt.Println("slice2的size=", len(slice2))
	// cap是一个内置函数,用于统计切片的容量,即最大可以存放多少个元素
	fmt.Println("slice2的cap=", cap(slice2))
	// 使用copy内置函数完成拷贝(切片拷贝如果空间小了,也不会报错)
	slice_copy := make([]int, 10)
	copy(slice_copy, slice2)
	fmt.Println("slice_copy = ", slice_copy)

	///// 切片的遍历
	var arr = [...]int{1, 2, 3, 4, 5}
	slice4 := arr[1:4]

	for i := 0; i < len(slice4); i++ {
		fmt.Printf("slice4[%v]=%v", i, slice4[i])
	}

	for k, v := range slice4 {
		fmt.Printf("slice4 %v = slice4[%v]=%v", k, slice4[k], v)
	}

	// 切片初始化方法写法
	slice5 := arr[0:3]
	slice5 = arr[0:] // 默认最大长度
	slice5 = arr[:]
	slice5 = arr[:3] // 默认最小策划长度

	// 切片可以继续切片
	slice6 := slice5[0:1]
	fmt.Println("slice6 = %v", slice6)

	// 切片可以动态增长
	var slice7 []int = []int{100, 200, 300}
	// 通过append直接给slice7 追加具体元素
	// 底层是通过值拷贝的形式,如果超过空间,会创建一个新的数组,并加新的值进行添加
	slice7 = append(slice7, 400)
	slice7 = append(slice7, 500, 600, 700)

	//// 切片slice和string
	// 切片的底层是byte数组
	str := "abcde@fgh"
	// 使用切片进行截取
	slice8 := str[6:]
	fmt.Println("slice8 = ", slice8)
	// str的内存类型数值不能改变
	// str[0] = 'a';

	// 如果需要修改字符串,则先将字符串转换为一个[]byte或者[]rune,重新转成一个string
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

}

// 切片的使用
func TestSliceFunc() {
	// 声明一个数组 切片
	arr := []int{1, 2, 3, 4, 5}
	// 切片长度
	fmt.Println("slice的size=", len(arr))
	// 切片容量
	fmt.Println("slice的cap=", cap(arr))
	var copy_arr = make([]int, 1)
	copy(copy_arr, arr)
	fmt.Println(copy_arr)

	// 切片遍历
	for k, v := range arr {
		fmt.Printf("key = %v , val = %v", k, v)
	}

	// 切片可以继续切片
	// 切片初始化方法
	newSlice := arr[0:3]
	newSlice = arr[0:] // 默认最大长度
	newSlice = arr[:]
	newSlice = arr[:3] // 默认最小策划长度

	fmt.Println("newSlice = ", newSlice)
	// 追加切片
	fmt.Println("old slice = ", newSlice)
	newSlice = append(newSlice, 101)
	fmt.Println("new slice = ", newSlice)

	// 如果字符串包含中文编码,则先将编码转换为rune,再进行修改
	str2 := "你好!"
	arr2 := []rune(str2)
	fmt.Println("len str2=", len(arr2))
	arr2[0] = '一'
	str2 = (string)(arr2)
	fmt.Println("str2=", str2)
}

//// 二维数组介绍,多维数组介绍
func testArrays() {
	// 定义声明一个二维数组
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 1

	fmt.Println(arr)

}
