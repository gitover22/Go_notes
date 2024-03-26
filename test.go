package main

import (
	"fmt"
)

/**
* @brife min(x^y,lim)
 */
func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21 // 通过指针修改指向元素的值
	fmt.Println(i)
	p = &j // p修改指向
	*p = *p / 37
	fmt.Println(j)

}
