/**
* @File  : zero.go
* @Author: Shide 2024/11/18 17:07
* @Desc  : 零值
 */

package main

import "fmt"

/*
*
没有明确初始化的变量声明会被赋予对应类型的 零值。

零值是：

数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）。
*/
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
