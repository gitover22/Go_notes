package main

import "fmt"

func add(x, y int) int { // 当连续两个或多个函数的已命名形参类型相同时，可省略形参
	return x + y
}

func main() {
	// rand.Seed(time.Now().UnixNano())
	// fmt.Printf("%s\n", "string output")
	// fmt.Println("random:", rand.Intn(100))
	// fmt.Println(math.Pi) // 大写字母开头的就是已导出的
	fmt.Println("add(10,20):", add(10, 20))
}
