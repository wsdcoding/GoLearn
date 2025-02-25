/**
* @File  : index.go
* @Author: Shide 2025/2/25 20:47
* @Desc  :
类型参数
可以使用类型参数编写 Go 函数来处理多种类型。 函数的类型参数出现在函数参数之前的方括号之间。

func Index[T comparable](s []T, x T) int
此声明意味着 s 是满足内置约束 comparable 的任何类型 T 的切片。 x 也是相同类型的值。

comparable 是一个有用的约束，它能让我们对任意满足该类型的值使用 == 和 != 运算符。在此示例中，我们使用它将值与所有切片元素进行比较，直到找到匹配项。 该 Index 函数适用于任何支持比较的类型。
*/

package main

import "fmt"

// Index 返回 x 在 s 中的下标，未找到则返回 -1。
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v 和 x 的类型为 T，它拥有 comparable 可比较的约束，
		// 因此我们可以使用 ==。
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index 可以在整数切片上使用
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index 也可以在字符串切片上使用
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
