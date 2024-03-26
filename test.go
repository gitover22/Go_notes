package main

import "fmt"

/**
* @brief mutil return values func test
*s
 */

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(split(17))
}
