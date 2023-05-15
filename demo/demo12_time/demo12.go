/// 日期与时间相关的函数

package demo

import (
	"fmt"
	"time"
)

// 日期时间
func testTime() {
	// 1.获取当前时间
	now := time.Now()
	fmt.Printf("now = %v type = %T", now, now)

	// 获取now的年月日,时分秒
	fmt.Printf("年 = %v \n", now.Year())
	fmt.Printf("月 = %v \n", now.Month())
	fmt.Printf("日 = %v \n", now.Day())
	fmt.Printf("时 = %v \n", now.Hour())
	fmt.Printf("分 = %v \n", now.Minute())
	fmt.Printf("秒 = %v \n", now.Second())

	// 格式化日期事件
	// 以yyyy-MM-dd HH:mm:ss格式输出日期时间字符串
	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// 格式化的时间字符串
	fmt.Printf("dateStr=%v\n", dateStr)

	// 使用now.Format 进行格式化
	// 必须使用 "2006-01-02 15:04:05" 字符串
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	// 时间常量
	// time.Duration 纳秒 1000纳秒等于1微秒
	// time.Microsecond //微秒 1000微秒等于1毫秒
	// time.Millisecond // 毫秒 1000毫秒等于1秒
	// time.Second // 秒
	// time.Minute // 分钟
	// time.Hour // 小时

	// 休眠
	time.Sleep(100 * time.Millisecond) // 休眠100毫秒

	// 获得unix 时间戳
	// 获取unixnano 时间戳
	fmt.Printf("unix 时间戳=%v unixnano时间戳=%v", now.Unix(), now.UnixNano())

}
