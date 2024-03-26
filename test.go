package main

import "fmt"

func main() {
	i := 10
	var f float32 = float32(i)
	var d float64 = float64(f)
	var b bool = bool(true)
	fmt.Printf("%v,  %v, %v, %v\n", i, f, d, b)
}
