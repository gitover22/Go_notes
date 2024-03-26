package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 值接收者只是在对副本进行操作
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 指针接收者能够改变值
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(v.X, v.Y)
	fmt.Println(Abs(v))
}
