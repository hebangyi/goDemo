package demo22

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	c := make(chan int)
	go Producer(c)
	go Consumer1(c)

	time.Sleep(time.Second * 20)
}

func Producer(c chan int) {
	for i := 0; i < 1; i++ {
		c <- i

		time.Sleep(time.Second * 1)
	}
	close(c)
	// 两次close 会panic
	// close(c)
}

func Consumer(c chan int) {
	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("Consumer Fin")
}

func Consumer1(c chan int) {
	for {
		select {
		case v, ok := <-c:
			fmt.Println("ok = ", ok)
			fmt.Println(v)
		}
	}
	fmt.Println("Consumer Fin1")
}
