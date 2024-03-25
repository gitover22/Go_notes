package main

import "fmt"

/**
* @brief mutil return values func test
*
 */

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	x := "world"
	y := "hello"
	z, u := swap(x, y)
	fmt.Printf("%s,%s", z, u)
}
