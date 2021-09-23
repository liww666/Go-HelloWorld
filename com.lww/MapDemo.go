package main

import (
	"fmt"
	"sort"
)

func main() {
	//声明map
	ages := make(map[string]int)
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)
	var sorces map[string]int
	fmt.Println(sorces == nil)
	fmt.Println(len(sorces) == 0)

	ages["alice"] = 23
	ages["charlie"] = 24
	fmt.Println(ages["charlie"]) //如果key存在，返回对应的value，否则返回value类型对应的零值
	delete(ages, "alice")        //删除指定key
	ages["charlie"]++
	fmt.Println(ages["charlie"])

	/**
	map中的元素并不是变量，不能对map元素进行取址操作
	因为map的元素地址可能随着元素数量的增大而重新分配更大的内存空间，
	从而导致之前的地址无效
	*/
	//var ptr  = &ages["charlie"]

	//声明同时赋初始值
	sorces = map[string]int{
		"a": 100,
		"b": 90,
		"c": 80,
		"d": 70,
		"e": 60,
	}
	fmt.Printf("a.sorce=%d\n", sorces["a"])
	fmt.Println("========无序遍历===============")
	foreach(sorces)
	fmt.Println("========有序遍历===============")
	foreachByName(sorces)

}

/**
map遍历
每一次的遍历顺序都不相同
*/
func foreach(sorces map[string]int) {
	for name, sorce := range sorces {
		fmt.Printf("%s\t%d\n", name, sorce)
	}
}

func foreachByName(sorces map[string]int) {
	names := make([]string, 0, len(sorces))
	for name := range sorces {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, sorces[name])
	}
}
