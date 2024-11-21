/**
* @File  : exercise_loops_and_functions.go
* @Author: Shide 2024/11/21 10:17
* @Desc  : 循环与函数的练习题，牛顿算法go语言实现
 */

package main

import (
	"fmt"
	"math"
)

// 方法一
func Sqrt1(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

// 方法2，调用函数无限循环
func Sqrt(x float64) float64 {
	z := float64(1)
	for {
		res := z - (z*z-x)/(2*z)
		if math.Abs(res-z) < 1e-6 {
			return res
		}
		z = res
	}
	return z
}
func main() {

	fmt.Println(Sqrt1(2))
	fmt.Println(math.Sqrt(9))
}
