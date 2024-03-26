package main

import (
	"fmt"
)

// 闭包的用法示例
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	} // 闭包，sum不会摧毁，会一直维护
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
