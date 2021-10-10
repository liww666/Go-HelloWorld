package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Connecting countdown")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
}
