/**
@File  : funcrtion_values.go
@Author: Shide 2024/11/21 15:15
@Desc  : 函数值
函数也是值。它们可以像其他值一样传递。
函数值可以用作函数的参数或返回值。

*/

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
