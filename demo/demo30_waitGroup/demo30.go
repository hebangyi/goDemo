package demo30

import (
	"fmt"
	"sync"
	"time"
)

func Test() {
	var wg sync.WaitGroup
	// 启动三个 Goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(5000)
			fmt.Printf("Goroutine %d 执行任务\n", id)
		}(i)
	}
	// 等待所有 Goroutine 完成任务
	wg.Wait()
	fmt.Println("所有任务已完成!")
}
