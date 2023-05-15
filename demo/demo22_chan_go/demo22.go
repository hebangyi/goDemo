package demo

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

/// 协程和管道和锁

// 主线程是一个物理线程,直接作用在CPU上,是由CPU调度,是重量级的,非常耗费cpu资源
// 协程是从主线程开启的,是轻量级的线程,是逻辑态
// 由gou语言调度 对资源消耗相对较小
// Golang的协程机制是重要的特点,可以轻松开启上万个协程,其它编程语言的并发机制一般基于线程的,开启过多的线程,资源耗费大,这里就凸显GoLang在并发上的优势了
// 协程采用MPG 模式
// M 核心线程
// P 上下文状态
// G 执行的协程

// 因为协程是逻辑态,所以可以很轻松的启动上万个协程

// 携程与管道
// 协程与多线程的区别
// 1.协程能像多线程一样,充分利用cpu资源
// 2.能再逻辑态上充分利用资源,减少开销,底层做了优化
// ? 协程的基本原理
// go中的协程更轻量,更稳健
// go协程的特点
// 1.有独立的栈空间
// 2.共享程序堆空间
// 3.调度由用户控制
// 4.协程是轻量级的线程

// 协程函数
func routine() {
	for i := 10; i < 10; i++ {
		fmt.Println("test () hello, world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// 测试协程函数
func testGoRoutine() {
	go routine() // 开启了一个协程
	// 输出效果
	for i := 10; i < 10; i++ {
		fmt.Println("main () hello, world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	// 如果主线程结束,那么整个程序也结束了
}

// 获得cpu的核数,并设置多核cpu
func getCpuNum() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	// 可以自己设置使用多个cpu
	// 在go1.8之前,使用多个cpu,需要设置
	// 1.8后,默认让程序运行在多核上,可以不用设置
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
	//  查看是否有数据竞争
	// go build -race main.go
}

var (
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁
	// sync 是包,
	lock sync.Mutex
)

func testLocK() {
	// 加锁
	lock.Lock()
	// 解锁
	defer lock.Unlock()
}

// channel 管道的基本介绍
// 1. channel 本质就是一个数据结构-队列
// 2. 数据是先进先出
// 3. 线程安全,多goroutine访问时,不需要加锁,就是说channel本身就是线程安全的
// 4. channel是有类型的,一个string的channel只能存放string类型数据
func testChannel() {

	// 演示一下管道的使用
	// 1. 创建一个可以存放3个int类型的管道,管道的容量不能自动增长
	var intChan chan int // 创建int类型的管道
	intChan = make(chan int, 3)

	// 2. 看看intChan是什么
	fmt.Printf("intChan 的值=%v\n", intChan)

	// 3. 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 注意点,当我们给管道写入数据是,不能超过其容量

	// 4.看看管道得长度和cap(容量)
	fmt.Printf("channel len = %v cap = %v\n", len(intChan), cap(intChan))

	// 5. 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2 = ", num2)
	fmt.Printf("channel len = %v cap = %v \n", len(intChan), cap(intChan))

	// 6.在没有使用协程的情况下,如果我们的管道全部取出,再取就会报告 deadlock
}

// channel 的关闭,channel关闭后仍然可以读取
// channel 的遍历
func testChannel2() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) // channel close
	// 内建函数close 关闭信道,该通道必须为双向的或只发送的,其效果是最后发送的值被接受后停止该通道,在最后的值从已关闭的信道中被接收后,任何对其接收操作都会无阻塞的成功

	// 这个时候不能再写入数到channel

	// channel 支持for-range 的方式进行遍历
	// 注意两个细节
	// 在遍历时,如果channel没有关闭,则会出现deadlock 错误,将阻塞
	// 在遍历时,如果channel已经关闭,则会正常遍历数据,遍历完后,就会退出遍历.

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 // 放入100个数据到管道
	}

	// 遍历管道时,不能使用普通的for循环结构
	// 需要使用for each
}

// 协程与通道结合相关,协程执行顺序
func testChannel3() {
	go writeData(channel3)
	go readData(channel3, channel3Exit)

	for {
		// 该方法将会阻塞,直到取完关闭
		v, ok := <-channel3Exit
		if !ok {
			break
		}
		// time.Sleep
		fmt.Printf("readData 读到数据 = %v \n", v)
	}
	// 结束
	fmt.Println("Exit!")
}

// 写入数据的协程
func writeData(intChan chan int) {

	// 这里有一个机制,注意,如果读取的数据很慢,channel3 超出所在的长度,也不会发生Data-lock,会写阻塞
	// 写管道和度管道频率不一致,无所谓
	for i := 1; i < 50; i++ {
		// 放入数据
		channel3 <- i
	}
	close(channel3)
}

func readData(intChan chan int, channel3Exit chan bool) {
	for {
		// 该方法将会阻塞,直到取完关闭
		v, ok := <-intChan
		if !ok {
			break
		}
		// time.Sleep
		fmt.Printf("readData 读到数据 = %v \n", v)
	}

	// 如果全部读取完成,向主协程发出信号
	channel3Exit <- true
	close(channel3Exit)
}

var (
	channel3     = make(chan int, 3)
	channel3Exit = make(chan bool, 1)
)

// /// 判断 1-8000是不是素数,使用协程的方式
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	// 关闭intChan
	close(intChan)
}

// 从intChan取出数据,并判断是否为素数,如果是,就放入到primeChan中
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	// 使用 for 循环
	var flag bool = true
	for {
		num, ok := <-intChan
		if !ok {
			break
		}

		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				// 说明该number不是素数
				flag = false
				break
			}
		}

		if flag {
			// 将这个数放入到 primeChan
			primeChan <- num
			//
		}
	}

	fmt.Println("有一个primeNum 协程因为取不到数据,退出")
	// 这里我们还不能关闭 primeChan
	exitChan <- true
}

func testPutNum() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) // 放入结果
	// 标识退出管道
	exitChan := make(chan bool, 4) // 4个

	// 开启一个协程,向 intChan 放入 1-8000个数
	go putNum(intChan)
	// 开启4个协程, 从 intChan 取出数据 , 并判断是否为素数 , 如果是, 就放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 使用协程退出标识
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
			close(primeChan)
		}
	}()

	// 主线程变量结果
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}

		// 将结果输出
		fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main thread exit")

}

func TestSelectCase() {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	intChan <- 1
	stringChan <- "hello"
	select {
	case value := <-intChan:
		fmt.Println(value)
	case value := <-stringChan:
		panic(value)
	}

}

func TestSingalChannel() {
	ch := make(chan int)
	go sendOnly(ch, 42)
	go receiveOnly(ch)
}

// 发送的单向通道
func sendOnly(ch chan<- int, value int) {
	ch <- value
}

// 接受的单向通道
func receiveOnly(ch <-chan int) {
	value := <-ch
	fmt.Println("Received:", value)
}
