package main

import (
	"fmt"
	"time"
)

/**
* @brife min(x^y,lim)
 */
func main() {
	t := time.Now()
	fmt.Println(t.Hour())
	switch { // 等价于switch true
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening")
	}

}
