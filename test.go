package main

import (
	"fmt"
	"math"
)

// 定义接口 I
type I interface {
	M() // 接口包含的方法
}

// 类型T
type T struct {
	S string
}

// 类型T的指针型方法M
func (t *T) M() {
	fmt.Println(t.S)
}

// 类型F
type F float64

// 类型F的值类型方法M
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
