package main

import "fmt"

type Veee struct {
	X int
	Y int
	Z *int
}

func main() {
	v := Veee{}
	v.X = 5
	v.Y = 8
	v.Z = &v.X
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(*v.Z)
}
