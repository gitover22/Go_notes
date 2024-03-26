package main

import "fmt"

func main() {
	// 创建长度为10的数组
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i + 1
	}
	fmt.Println(a)
	// 创建的同时初试化
	prime := [5]int{5, 4, 3, 2, 1}
	fmt.Println(prime)
	// slices
	var aa []int = a[5:10]
	fmt.Println(aa)
}
