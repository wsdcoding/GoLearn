/**
* @File  : numeric_constants.go
* @Author: Shide 2024/11/18 17:26
* @Desc  : 数值常量
 */

package main

import "fmt"

/*
*
数值常量是高精度的 值。

一个未指定类型的常量由上下文来决定其类型。

再试着一下输出 needInt(Big) 吧。

（int 类型可以存储最大 64 位的整数，根据平台不同有时会更小。）
*/
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
