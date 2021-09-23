package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main(){
	s:=[]int{1,3,4,6}
	c:= make(chan int)
	//d:=make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:],c)


	x ,y:=<-c,<-c
	fmt.Println(x,y,x+y)
}