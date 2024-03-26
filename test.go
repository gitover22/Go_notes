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
	} else {
		fmt.Printf("%g >= %g\n", v, lim) //%g能自动选择合适的输出格式
	}
	return lim
}
func main() {
	fmt.Println(__min(3, 2, 10), __min(3, 3, 10))
}
