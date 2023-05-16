package demo29

import "fmt"

func Test1() {
	c := &Context{StartNum: 10, Nums: make([]int, 0), Step: 0}
	Calc(c, c.StartNum)
}

func Test2() {
	go Write()
	for num := 0; num < 20; num++ {
		go Execute()
	}
}

var c chan int = make(chan int, 10000)

func Write() {
	num := 0
	for {
		c <- num
		num++
	}
}

func Execute() {
	num := 0
	context := &Context{StartNum: num, Nums: make([]int, 0), Step: 0}
	fmt.Println("进入了一个协程!!!")
	for {
		num = <-c
		Calc(context, context.StartNum)
		context.StartNum = num
		context.Step = 0
		context.Nums = context.Nums[:0]
	}
}

type Context struct {
	StartNum int
	Nums     []int
	Step     int
}

func (c *Context) Contains(num int) bool {
	for newNum := range c.Nums {
		if newNum == num {
			return true
		}
	}
	return false
}

func (c *Context) Show() {
	fmt.Printf("StartNum = %v  Nums = %v , Step = %v\n", c.StartNum, c.Nums, c.Step)
}

func Calc(c *Context, num int) {
	if c.Contains(num) {
		// c.Show()
		if c.StartNum%10000 == 0 {
			c.Show()
		}
		return
	}

	if num%2 == 0 {
		num = num / 2
	} else {
		num = num*3 + 1
	}
	c.Step++
	c.Nums = append(c.Nums, num)
	Calc(c, num)
}
