package main

import "fmt"

func main() {
	// 切片默认为从0到尾
	num := []int{1, 2, 3, 4, 5}
	aa := num[:]
	fmt.Println(aa)
}
