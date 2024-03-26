package main

import (
	"fmt"
	"runtime"
)

/**
* @brife min(x^y,lim)
 */
func main() {
	fmt.Print("go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")

	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s\n", os)
	}
}
