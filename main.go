package main

import (
	"flag"
)

var path string
var method string

func init() {
	flag.StringVar(&path, "path", "aa", "aa")
	flag.StringVar(&path, "method", "aa", "aa")
}

func main() {
	/*	flag.Parse()
		fmt.Println("soPath: ", path)
		fmt.Println("method: ", method)
		ptr, err := plugin.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		methodFunc, _ := ptr.Lookup(method)
		methodFunc.(func())()*/

	/*	Add, _ := ptr.Lookup("Add")
		sum := Add.(func(int, int) int)(5, 4)
		fmt.Println("Add结果：", sum)

		Sub, _ := ptr.Lookup("Subtract")
		sub := Sub.(func(int, int) int)(9, 8)
		fmt.Println("Sub结果：", sub)*/

	//cmd := exec.Command("demo\\demo.exe")
	//output, err := cmd.CombinedOutput()
	//fmt.Printf("output: %s\n", string(output))
	//fmt.Printf("err: %v", err)

	//var c1 = make(chan int)
	//var c2 = make(chan int)
	//
	//select {
	//case <-c1:
	//	fmt.Println("c1")
	//case <-c2:
	//	fmt.Println("c2")
	//default:
	//	fmt.Println("default")
	//}
	//fmt.Println("main")
	select {}

}
