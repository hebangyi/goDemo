package demo

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串相关
func testStr1() {
	var str1 string = "hello"
	// 统计字符串的字节长度, 按字节len(str)
	//// golang的编码格式为为utf-8 (ascii的字符(字母和数值) 占一个字节,汉字占用3个字节)
	fmt.Println("str len=", len(str1))
	// rune 本质是int32 , 用来表示unicode 字符
	// 所以在计算字符串长度时 应该先通过rune函数进行转换
	str2 := []rune("hello 你好")
	for i := 0; i < len(str2); i++ {
		fmt.Printf("")
	}

	// 字符串转为整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("字符串转换错误!")
	} else {
		fmt.Println("转换的结果是:", n)
	}

	// 整数转字符串
	str3 := strconv.Itoa(123456)
	fmt.Printf("str=%v,str=%T", str3, str3)

	// 字符串转 []byte
	bytes := []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	// []byte 转字符串
	str4 := string([]byte{97, 98, 99})
	fmt.Println("str4 = ", str4)

	// 10 进制 转为 对应的 2 8 16 进制
	str5 := strconv.FormatInt(123, 2)
	fmt.Println("123对应的二进制是=%v", str5)
	str5 = strconv.FormatInt(123, 16)
	fmt.Println("123对应的十六进制是=%v", str5)

	// 查找子串是否在指定的字符串中
	b := strings.Contains("seafood", "foo") // true
	fmt.Printf("b = %v", b)

	// 统计字符串中有几个子串
	count := strings.Count("abcba", "a")
	fmt.Printf("count = %v", count)

	// 在比较字符串 == 与是区分大小写的 , 判断字符串相同,不区分大小写
	b = strings.EqualFold("abc", "ABC")
	fmt.Printf("b = %v", b)
	fmt.Printf("b = %v", "abc" == "ABC")
	// 返回子串在字符串第一次出现的index值 (从0开始计算) , 如果没有,则返回-1
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("index 出现的位置 = %v", index)
	// 返回子串最后一次出现的index
	lastIndex := strings.LastIndex("NLT_bbbbc", "b")
	fmt.Printf("lastIndex 最后出现的位置 = %v", lastIndex)
	// 字符串替换,如果num = -1 表示全部替换
	replaceNum := 1
	replaceStr := strings.Replace("gogogo", "go", "go语言", replaceNum)
	fmt.Printf("替换的字符串为v = %v", replaceStr)
	// 字符串按照指定的字符分割
	strArr := strings.Split("hello,wrold,go", ",")
	fmt.Println("strArr Len = ", len(strArr))
	// 将字符串进行大小写转换
	fmt.Printf("小写字母 = %v", strings.ToLower("Go"))
	fmt.Printf("大写字母 = %v", strings.ToLower("Go"))
	// 字符串去左右空格
	fmt.Printf("去掉空格的字符串 = %v", strings.TrimSpace(" sdd dddd "))
	// 字符串去指定的字符串
	fmt.Printf("去除指定的字符串 = %v", strings.Trim(" Go hello", "Go"))
	// 将字符串左边指定的字符去掉
	fmt.Printf("去除指定的字符串 = %v", strings.TrimLeft(" Go hello Go", "Go"))
	// 将字符串右边指定的字符去掉
	fmt.Printf("去除指定的字符串 = %v", strings.TrimRight(" Go hello Go", "Go"))

	fmt.Printf("字符串是否以 XXX开头 = %v", strings.HasPrefix("XXXABCD", "XXX"))
	fmt.Printf("字符串是否以 XXX结束 = %v", strings.HasSuffix("XXXABCD", "BCD"))
}
