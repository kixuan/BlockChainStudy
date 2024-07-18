package main

import (
	"fmt"
	"time"
)

func work(ch chan int) {
	time.Sleep(5 * time.Second)
	ch <- 5
}

func main() {
	ch := make(chan int, 10)
	now_time := time.Now()
	ch <- 5
	fmt.Printf("当前时间 %d-%d-%d %d:%d:%d \r\n", now_time.Year(), int(now_time.Month()), now_time.Day(), now_time.Hour(), now_time.Minute(), now_time.Second())
	go work(ch)                 //启动协程调用work()
	time.Sleep(5 * time.Second) //模拟做其他事情，假定耗时5s
	sec := <-ch
	fmt.Printf("调用work()用时%ds.\r\n", sec)
	now_time = time.Now()
	fmt.Printf("当前时间 %d-%d-%d %d:%d:%d \n", now_time.Year(), int(now_time.Month()), now_time.Day(), now_time.Hour(), now_time.Minute(), now_time.Second())
	close(ch)

}
