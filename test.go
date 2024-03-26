package main

import "fmt"

// len():返回切片元素个数
// cap():从切片第一个元素开始数，到其底层数组元素末尾的个数
func main() {
	// 切片默认为从0到尾
	num := []int{1, 2, 3, 4, 5}
	aa := num[1:2]
	printSlice(aa)
}
func printSlice(s []int) {
	fmt.Printf("len: %d , casp:%d  v: %v\n", len(s), cap(s), s)
}
