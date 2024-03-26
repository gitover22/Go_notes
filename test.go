package main

import "fmt"

/**
*
*
 */
var i, j int = 2, 3

// k := 3 函数外不能这样使用，必须使用var func
func main() {
	var c, python, java = true, false, "string" // 默认为flase
	k := "fsdaf"
	fmt.Println(i, j, c, python, java, k)
}
