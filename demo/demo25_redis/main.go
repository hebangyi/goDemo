package main

/// 第三方包导入
/// redis 朝着
/// go get 导入第三方表
import (
	"fmt"
	// 还有一个封装redis 具体方法的包
	"github.com/gomodule/redigo/redis"
)

// redis 练习
func main() {
	// 使用 go get 命令在 gopath 下导入第三方开发包
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Redis Client Error !")
	}
	defer conn.Close()

	_, err = conn.Do("set", "name", "heby")
	if err != nil {
		fmt.Println("set err = ", err)
		return
	}

	fmt.Println("redis set 完成!")
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Printf("取值错误 : err = %v", err)
	}

	// 因为返回的r是interface{}.因此我们需要转换
	fmt.Println("result = ", r)
	// 设置name的数据有效时间为10s
	conn.Do("expire", "name", 10)

}

func TestRedisPool() {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			// 初始化代码
			return redis.Dial("tcp", "localhost:6379")
		},
		MaxIdle:     8,   // 最大空闲链接数
		MaxActive:   0,   // 最大和数据库链接数 0表示没有限制
		IdleTimeout: 100, // 最大空闲链接时间
	}

	// 从redis 中获得一个链接
	conn := pool.Get()
	// 将连接返回到连接池
	defer conn.Close()
	conn.Do("set", "name", "tom")
}
