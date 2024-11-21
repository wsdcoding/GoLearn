/**
@File  : array.go
@Author: Shide 2024/11/21 11:10
@Desc  : 数组
类型 [n]T 表示一个数组，它拥有 n 个类型为 T 的值。
表达式
var a [10]int
会将变量 a 声明为拥有 10 个整数的数组。
数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是个限制，不过没关系，Go 拥有更加方便的使用数组的方式。

*/

package main

import "fmt"

func main() {
	var s [10]string

	s[0] = "a"
	s[1] = "b"
	fmt.Println(s[0], s[0])
	fmt.Println(s)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
