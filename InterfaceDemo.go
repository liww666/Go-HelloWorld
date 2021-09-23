package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
	
}
type IPhone struct {

}

func (nokiaPhone NokiaPhone)  call() {
	fmt.Println("I am nokia,i can call you")
}

func (iphone IPhone) call()  {
	fmt.Println("I am iphone,i can call you")
}

func main()  {
	var phone Phone
	phone = new (NokiaPhone)
	phone.call()

	phone = new (IPhone)
	phone.call()
}