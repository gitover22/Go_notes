package main

import "fmt"

const PI = 3.1415
const i = 3

func main() {
	const ss = "const string"
	const f = 3.23
	fmt.Printf("type: %T,  %T, %T, %T\n", PI, i, ss, f)
	fmt.Println("hello", PI)
	fmt.Println("hello", ss)
}
