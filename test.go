package main

import (
	"fmt"
	"time"
)

/**
* @brife min(x^y,lim)
 */
func main() {
	fmt.Print("when is Saturday")
	today := time.Now().Weekday()
	// fmt.Println(today + 2)
	switch time.Thursday {
	case today + 0:
		fmt.Println("today is Saturday")
	case today + 1:
		fmt.Println("tomorrow is Saturday")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("too far")
	}

}
