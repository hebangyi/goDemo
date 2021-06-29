/// GO Socket
package main

import (
	"fmt"
	"net" // 做网络 socket 开发时,net包含有我们需要所有的方法和函数
)

func main() {
	fmt.Println("服务器开始监听....")
	// 1. tcp 表示网络网络协议是tcp
	// 2. 0.0.0.0 表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}

	fmt.Printf("listen suc = %v", listen)
	//

	// 延时关闭
	defer listen.Close()
	// 循环等待客户端连接我

	for {
		// 等待客户端来链接
		fmt.Println("等待客户端来链接")
		// 获取连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept Error", err)
			continue
		} else {
			fmt.Printf("Accept Success ! conn = %v\n , 客户端IP= %v", conn, conn.RemoteAddr())
		}

		// 这里准备一个协程,为客户端服务
		go process(conn)
	}
}

func process(conn net.Conn) {
	// 这里我们循环接受客户端发送的数据
	defer conn.Close() // 关闭conn

	for {
		// 光剑一个新的切片
		buf := make([]byte, 1024)
		// 从conn读取
		// 1. 等待客户端通过 conn 发送信息
		// 2. 如果客户端没有 write[发送], 那么协程将阻塞
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的 Read err=", err)
			return
		}

		// 3. 显示客户端发送的内容到服务器的终端
		fmt.Printf(string(buf[:n]))
	}

}
