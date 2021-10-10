package main

import (
	"fmt"
	"sort"
	"strings"

	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}

	/*
		函数调用函数,声明函数变量
		函数值可以与nil比较，但是函数值之间是不能比较的，也不能作为map的key
	*/
	f := square
	fmt.Println(f(3))

	f1 := squares()
	fmt.Println(f1())

	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	/**
	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d // 必须把循环变量赋值给新的局部变量 NOTE: necessary!
		os.MkdirAll(dir, 0755) // creates parent directories too
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
	*/

	/**
	在上面的代码中，调用者隐式的创建一个数组，并将原始参数复制到数组中，再把数组的一
	个切片作为参数传给被调函数。如果原始参数已经是切片类型，我们该如何传递给sum？只需
	在最后一个参数后加上省略符。
	*/
	fmt.Println(sum(3))
	vals := []int{1, 2, 3}
	fmt.Println(sum(vals...))

	fmt.Println(double(2))
	fmt.Println(triple(2))
}

/**
多返回值函数
*/
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)
	}
	//doc, err := template.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML:%v", url, err)
	}
	//return visit(nil, doc), nil
	return nil, nil
}

func square(n int) int {
	return n * n
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

/**
可变参数函数
虽然在可变参数函数内部，...int 型参数的行为看起来很像切片类型，但实际上，可变参数函
数和以切片作为参数的函数是不同的。
*/
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	/*
		当defer语
		句被执行时，跟在defer后面的函数会被延迟执行。直到包含该defer语句的函数执行完毕时，
		defer后的函数才会被执行，不论包含defer语句的函数是通过return正常结束，还是由于panic
		导致的异常结束。你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相
		反。
	*/
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	//doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	// ...print doc's title element…
	return nil
}

func double(x int) (result int) {
	if x > 10 {
		return
	}
	y := x + x
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return double(y)
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

/**
错误的关闭文件方式
*/
func processFiles(filenames []string) error {
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		//f.close()只有在filenames遍历结束后才执行
		defer f.Close() // NOTE: risky; could run out of file
		// ...process f…
		return nil
	}
	return nil
}

/**
正确的关闭文件方式，
将循环体中的defer语句移至另外一个函数。在每次循环时，调用这个函数
*/
func proFiles(filenames []string) error {
	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			return err
		}
	}
	return nil
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...process f…
	return nil
}
