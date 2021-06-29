package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	fmt.Println("conn 成功 = ", conn)

	reader := bufio.NewReader(os.Stdin) // os.Stdin 代表标准输入[终端]

	// 重终端读取一行用户的输入并发送到服务器
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readString err = ", err)
	}

	// 再将 line 发送给服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn,Write err=", err)
	}
	fmt.Printf("客户端发送了 %d 字节的数据, 并退出", n)
}
