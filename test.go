package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "aa"
	ch <- "bb"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
