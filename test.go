package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (y Vertex) Abs() float64 {
	return math.Sqrt(y.X*y.X + y.Y*y.Y)
} // 结构体的函数
func Abs(x Vertex) float64 {
	return math.Sqrt(x.X*x.X + x.Y*x.Y)
} // 普通函数
func main() {
	fsdada := Vertex{3, 4}
	fmt.Println(fsdada.Abs(), Abs(fsdada))
}
