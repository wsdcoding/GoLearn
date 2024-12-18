/**
@File  : making_slices.go
@Author: Shide 2024/11/21 11:32
@Desc  : 用 make 创建切片
切片可以用内置函数 make 来创建，这也是你创建动态数组的方式。
make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
a := make([]int, 5)  // len(a)=5
要指定它的容量，需向 make 传入第三个参数：
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4

*/

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice_making_slices("a", a)

	b := make([]int, 0, 5)
	printSlice_making_slices("b", b)

	c := b[:2]
	printSlice_making_slices("c", c)

	d := c[2:5]
	printSlice_making_slices("d", d)
}

func printSlice_making_slices(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
