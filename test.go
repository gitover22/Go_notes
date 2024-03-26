package main

import (
	"fmt"
	"math"
)

/**
* @brife min(x^y,lim)
 */
func __min(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim { // if后是可以加一个执行语句的
		return v
	}
	return lim
}
func main() {
	fmt.Println(__min(3, 2, 10), __min(3, 3, 10))
}
