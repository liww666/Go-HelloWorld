package main

import "fmt"

var t1,t2 int //全局变量可以声明不赋值和使用
//t3:=4 全局变量不可这样声明和赋值
func main() {
	fmt.Println("Hello World")
	plus()
}

func plus() {
	var a, b string = "a", "b"
	var c string
	fmt.Println(a + b + "c= " + c)


	var i int
	var f float64
	var e bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, e, s)

	var d = true
	fmt.Println(d)
	g := false //声明语句，类型 var g = true
	fmt.Println(g)

	var v1, v2, v3 = 1, 2, 3
	var v4 = v1
	fmt.Printf("%v,%v,%v\n", v1, v2, v3)
	fmt.Printf("%v\n",&v1)
	fmt.Printf("%v\n",&v2)
	fmt.Printf("%v\n",&v3)
	fmt.Printf("%v\n",&v4)
	var l1 int
	fmt.Println("t1="+string(t1))
	fmt.Println("l1="+string(l1))

	const c1 = 1
	fmt.Println("c1"+string(c1))
	
	var ptr *int = &v1
	fmt.Println(*ptr)
	x,y:=swap("a","b")
	fmt.Println(x,y)

	var balance = [5]float32{11,1,1,1,1}
	fmt.Println(balance)

	fmt.Println(ptr)

	var p1 *string
	fmt.Println(p1)
	if(p1==nil){
		fmt.Println(p1)
	}else{
		fmt.Println(*p1)
	}

	var a1,a2 string = "1","2"
	swap2(&a1,&a2)
	fmt.Println(string(a1)+","+string(a2))

	fmt.Println(Book{"go语言","lww","计算机"})
	var book = Book{"java","lww","jsj"}
	var bookPtr *Book = &book
	fmt.Println(book.title)
	fmt.Println(&book.title)
	fmt.Println(bookPtr.author)

}

func swap(x,y string)(string,string){
	return y,x
}

func swap2(x *string,y *string){
	var temp = *x
	*x = *y
	*y = temp
}

type Book struct {
	title string
	author string
	subject string
}