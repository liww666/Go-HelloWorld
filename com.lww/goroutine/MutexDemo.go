package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter1 int
	wg2      sync.WaitGroup
	mutex    sync.Mutex
)

func main() {
	wg2.Add(2)
	go incCounter1(1)
	go incCounter1(2)
	wg2.Wait()
	fmt.Printf("Final counter1:%d\n", counter1)

}

func incCounter1(id int) {
	defer wg2.Done()
	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter1
			runtime.Gosched()
			value++
			counter1 = value
		}
		mutex.Unlock()
	}
}
