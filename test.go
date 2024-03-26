package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8) // b为数组切片，len为8，cap为8
	for {
		n, err := r.Read(b) //填充b，并返回填充的字节数和错误值
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF { // 遇到结尾，返回EOF
			break
		}
	}
}
