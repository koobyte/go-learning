package main

import (
	"fmt"
	"time"
)

func main() {
	// println("===== showLetters =====")
	// goroutine.ShowLetters()
	// println("===== PrintPrimeNumber =====")
	// goroutine.PrintPrimeNumber()
	// println("===== Demo =====")
	// goroutine.Demo()
	// println("===== Demo1 =====")
	// goroutine.Demo1()

	println("===== 竞态条件 =====")

	// 存在竞态条件
	// 可以进入main包，执行 go build -race main.go，此时可以编译出竞态条件的可执行文件，然后执行 ./main，运行会出错：
	// WARNING: DATA RACE
	// Read at 0x0000011f4920 by goroutine 8:
	//  github.com/huzhouv/go-learning/goroutine.IncErr()
	// ......
	// 成功检测到存在竞态条件的代码而执行失败
	// goroutine.Count()
	// goroutine.Countn()
	// goroutine.SyncCountn()
	// goroutine.AtomicCountn()
	// goroutine.BreakFor()

	println("===== chan =====")
	// goroutine.PingPong()
	// goroutine.Runner()
	// goroutine.MultiWork()

	println("======")

	// 使用 select 实现随机向channel写入0、1
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
			time.Sleep(time.Second)
		case ch <- 1:
			time.Sleep(time.Second)
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}
}
