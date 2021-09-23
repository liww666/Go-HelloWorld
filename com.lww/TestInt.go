package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	q := new(int)
	fmt.Println(q == p)

	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i,i+1,i*i)

	fmt.Printf("%08b\n", u)

	arr:= [...]int{1,2,3}
	fmt.Println(arr)

	months :=[...]string{1:"一",2:"二",3:"三",4:"四",5:"五",6:"六",7:"七",8:"八",9:"九",10:"十",11:"十一",12:"十二"}

	Q2 := months[4:7]
	summer :=months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	for _,s:=range summer{
		for _,q := range Q2{
			if s==q {
				fmt.Printf("%s appears in both\n ",s)
			}
		}
	}

	//fmt.Println(summer[:20])//panic:out of range

	endlessSummer :=summer[:5]
	fmt.Println(endlessSummer)
	reverse(months[:])
	fmt.Println(months)

	arr6:=[...]int{1,2,3,4,5,6}
	reverseInt(arr6[:3])
	fmt.Println(arr6)
	reverseInt(arr6[3:])
	fmt.Println(arr6)
	reverseInt(arr6[:])
	fmt.Println(arr6)
	mySlice:= make([]int,3,5)
	fmt.Println("mySlice:",mySlice)

	var s1 []int
	s1 =nil
	fmt.Println(s1==nil)
	s1 = []int(nil)
	fmt.Println(s1==nil)
	s1 =[]int{}
	fmt.Println(s1==nil)

	var nil_s []int
	empty_s:= []int{}
	println(nil_s)
	println(empty_s)

	str:= "Hello, 世界"
	fmt.Println("len(s):",len(str))//输出的是字节数，因为go使用的是utf-8编码，一个汉字占3个字节
	fmt.Println("rune:",len([]rune(str)))//输出的字符数
	var runes []rune
	for _,r:=range str{
		runes=append(runes,r)
	}
	fmt.Printf("%q\n",runes)

	testAppend()
	testAppend2()

}

func reverse(s []string)  {
	for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1{
		s[i],s[j] = s[j],s[i]
	}
	
}

func reverseInt(s []int)  {
	for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1{
		s[i],s[j] = s[j],s[i]
	}

}

func testAppend()  {
	mySlice:=[]int{1,2,3,4,5,6}
	/**
	当底层数组需要扩容时，会按照当前底层数组长度的2倍进行扩容，并生成新数组。
	如果底层数组的长度超过1000时，将按照25%的比率扩容，也就是1000个元素时，将扩展为1250个，不过这个增长比率的算法可能会随着go版本的递进而改变。
	 */
	newSlice:= append(mySlice,7)

	mySlice[3]= 33
	fmt.Println("old slice:",mySlice)
	fmt.Println("new slice:",newSlice)

	fmt.Println(len(mySlice),":",cap(mySlice))
	fmt.Println(len(newSlice),":",cap(newSlice))
	
}

func testAppend2()  {
	mySlice:=[]int{1,2,3,4,5,6}
	newSlice := mySlice[1:3:3]
	appSlice:= append(newSlice,7)

	mySlice[2]= 33
	fmt.Println("old slice:",mySlice)
	fmt.Println("new slice:",newSlice)
	fmt.Println("app slice:",appSlice)

	fmt.Println(len(mySlice),":",cap(mySlice))
	fmt.Println(len(newSlice),":",cap(newSlice))
	fmt.Println(len(appSlice),":",cap(appSlice))

}
