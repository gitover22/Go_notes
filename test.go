package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64}

func main() {
	num := []int{1, 3, 5, 3, 6}
	fmt.Println(num)
	// range 返回下标和值
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
