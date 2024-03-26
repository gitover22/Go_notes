package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	// 切片的零值是nil,且此时切片没有底层数组
	if s == nil {
		fmt.Println("nil!")
	}
}
