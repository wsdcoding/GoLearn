/**
* @File  : if_and else.go
* @Author: Shide 2024/11/21 10:12
* @Desc  : 在 if 的简短语句中声明的变量同样可以在对应的任何 else 块中使用。
（在 main 的 fmt.Println 调用开始前，两次对 pow 的调用均已执行并返回其各自的结果。）
*/

package main

import (
	"fmt"
	"math"
)

func a(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		a(3, 2, 10),
		a(3, 3, 20),
	)
}
