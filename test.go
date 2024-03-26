package main

import (
	"fmt"
)

type I interface {
	M() // 接口包含的方法
}

// 类型T
type T struct {
	S string
}

// 类型只要实现了一个接口的全部方法就说明实现了该接口
// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"} // 因为T类型实现了该接口，所以可以这样直接赋值
	i.M()
}
