package main

import "fmt"

func main() {
	// 自动类型推导
	i := 10
	f := 3.1415926
	c := 1 + 2i
	t := false
	// 显示说明
	var ii int = 12
	fmt.Printf("type: %T,  %T, %T, %T,%T\n", i, f, c, t, ii)
}
