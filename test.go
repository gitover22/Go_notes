package main

import "fmt"

func main() {
	names := []string{"zou", "gan"}
	fmt.Println(names)
	s := []struct {
		x int
		y bool
	}{
		{1, true},
		{0, false},
	}
	fmt.Println(s)
}
