package main

import (
	"fmt"
	"math/cmplx"
)

var (
	tobe    bool       = true
	MAX_INT uint64     = 1<<64 - 1
	z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("type: %T ,value:%v\n", tobe, tobe)
	fmt.Printf("type: %T ,value:%v\n", MAX_INT, MAX_INT)
	fmt.Printf("type: %T ,value:%v\n", z, z)

}
