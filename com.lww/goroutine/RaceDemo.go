package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	cnter int64
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	//go incCounter(1)
	//go incCounter(2)
	//原子新增
	go incCounterAtomic(1)
	go incCounterAtomic(2)
	wg.Wait()
	fmt.Println("Final cnter:", cnter)

}
func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := cnter
		runtime.Gosched()
		value++
		cnter = value
	}
}

func incCounterAtomic(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&cnter, 1)
		runtime.Gosched()

	}
}
