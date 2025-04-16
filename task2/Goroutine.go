package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型，为一个无参数无返回值的函数
type Task func()

// ScheduleTasks 函数用于调度并执行任务，同时统计每个任务的执行时间
func ScheduleTasks(tasks []Task) {
	var wg sync.WaitGroup
	taskTimes := make(map[int]time.Duration)

	for i, task := range tasks {
		wg.Add(1)
		go func(index int, t Task) {
			defer wg.Done()
			start := time.Now()
			t()
			elapsed := time.Since(start)
			taskTimes[index] = elapsed
		}(i, task)
	}

	wg.Wait()

	// 输出每个任务的执行时间
	for i, duration := range taskTimes {
		fmt.Printf("任务 %d 执行时间: %v\n", i+1, duration)
	}
}

func main() {
	// 	题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	// 考察点 ： go 关键字的使用、协程的并发执行。
	var wg sync.WaitGroup
	wg.Add(2)

	// 打印奇数
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("奇数:", i)
		}
	}()

	// 打印偶数
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("偶数:", i)
		}
	}()

	wg.Wait()
	// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	// 考察点 ：协程原理、并发任务调度。
	// 定义一组任务
	tasks := []Task{
		func() {
			time.Sleep(2 * time.Second)
		},
		func() {
			time.Sleep(1 * time.Second)
		},
		func() {
			time.Sleep(3 * time.Second)
		},
	}

	// 调度并执行任务
	ScheduleTasks(tasks)
}
