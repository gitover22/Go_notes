package main

import (
	"fmt"
)

/**
* @brife min(x^y,lim)
 */
func main() {
	fmt.Println("begin")
	defer fmt.Println("end")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //压栈
	}

}
