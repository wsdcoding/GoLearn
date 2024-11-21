/**
@File  : slices_pointers.go
@Author: Shide 2024/11/21 11:19
@Desc  : 切片类似数组的引用
切片就像数组的引用 切片并不存储任何数据，它只是描述了底层数组中的一段。
更改切片的元素会修改其底层数组中对应的元素。
和它共享底层数组的切片都会观测到这些修改。

*/

package main

import "fmt"

func main() {
	names := [4]string{
		"John0",
		"Paul1",
		"George2",
		"Ringo3",
	}
	fmt.Println(names) //[John0 Paul1 George2 Ringo3]
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) //[John0 Paul1] [Paul1 George2]

	b[0] = "XX"
	fmt.Println(a, b)  //[John0 XX] [XX George2]
	fmt.Println(names) //[John0 XX George2 Ringo3]

}
