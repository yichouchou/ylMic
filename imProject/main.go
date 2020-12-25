package main

import (
	"fmt"
	"sync"
	"time"
)

var i int64 = 10
var wg sync.WaitGroup

func SayHello(c chan int64) {
	defer wg.Done()
	i += 1
	c <- i //写入数据
	fmt.Println("hello, I am ", i)
}
func SayGoodbye(c chan int64) {
	defer wg.Done()
	fmt.Println("goodbye ", <-c) //读取数据
}
func main2() {
	wg.Add(2)
	c := make(chan int64)
	//通过c的通信完成数据共享
	go SayHello(c)
	go SayGoodbye(c)
	wg.Wait()
	fmt.Println("all done")

	ch1 := make(chan int, 4)
	ch2 := make(chan int, 4)
	ch3 := make(chan int, 4)
	ch3 <- 3
	select {
	case ch1 <- 1:
		fmt.Println("case 1")
		fmt.Println("ch1 is ", <-ch1)
	case ch2 <- 2:
		fmt.Println("case 2")
		fmt.Println("ch2 is ", <-ch2)
	case <-ch3:
		fmt.Println("case 3")
	case <-time.After(time.Second * 1):
		fmt.Println("case timeout")
	default:
		fmt.Println("case default")
	}
	select {
	case ch2 <- 5:
		fmt.Println("final", <-ch2)
	}
}

func main() {
	ints := make(chan int, 500)

	go func() {
		for i := 0; i < 100; i++ {
			ints <- 1
			fmt.Println("总量：--", len(ints))
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			<-ints
			fmt.Println("取走一瓶", len(ints))
		}

	}()

	select {
	case <-time.After(time.Second * 10):
		fmt.Println("case timeout")
	}
}
