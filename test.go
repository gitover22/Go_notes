package main

import "fmt"

func main() {
	names := [5]string{
		"111",
		"222",
		"333",
		"444",
		"555",
	}
	fmt.Println(names)
	var ss []string = names[0:2]
	var bb []string = names[1:3]
	fmt.Println(ss)
	fmt.Println(bb)

	ss[1] = "666" // 通过切片修改值，会影响原值
	fmt.Println("after change")
	fmt.Println(ss)
	fmt.Println(bb)
	fmt.Println(names)
}
