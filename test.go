package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 播种
	fmt.Printf("%s\n", "string output")
	fmt.Println("random:", rand.Intn(100)) // 取随机数[0,100)
	fmt.Println(math.Pi)                   // 大写字母开头的就是已导出的
}
