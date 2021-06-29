package demo

import (
	"flag"
	"fmt"
)

/// 方法,打印参数
func getName() string {
	return "Heby"
}

// 获得命令行参数
func testOsArgs() {
	// args := os.Args

}

func testFlage() {
	var user string
	var pwd string
	var host string
	var port int

	// &user 就是接受用户命令行中输入的 -u 后面的参数值
	// "u" 就是 -u 指定参数
	// "" 默认值
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名localhost")
	flag.IntVar(&port, "port", 3306, "端口号,默认为3306")

	// 这里有一个非常重要的操作,转换 , 必须调用该方法
	flag.Parse()
	// 输出结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
