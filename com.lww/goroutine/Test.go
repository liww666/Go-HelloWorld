package main

import (
	"fmt"
	"time"
)

func main() {
	chal := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		chal <- 1
	}()
	<-chal

	fmt.Println(cap(chal))

}
