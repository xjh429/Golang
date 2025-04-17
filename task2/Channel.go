package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	// 考察点 ：通道的基本使用、协程间通信。
	// 无缓冲通道
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println(num)
		}
	}()
	// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	// 考察点 ：通道的缓冲机制。
	// 有缓冲通道
	ch1 := make(chan int, 100)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 100; j++ {
			ch1 <- j
		}
		close(ch1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch1 {
			fmt.Println(num)
		}
	}()

	// 等待所有协程完成
	wg.Wait()
}
