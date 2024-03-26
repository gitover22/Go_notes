package main

import (
	"fmt"
)

// 返回一个“返回int的函数”
func fibonacci() func() int {
	a, b := 0, 1 // 斐波那契数列的前两个数
	return func() int {
		// 计算下一个斐波那契数
		next := a
		a, b = b, a+b
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
