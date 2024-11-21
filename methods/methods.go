/**
@File  : methods.go
@Author: Shide 2024/11/21 15:31
@Desc  : 方法
Go 没有类。不过你可以为类型定义方法。
方法就是一类带特殊的 接收者 参数的函数。
方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
在此例中，Abs 方法拥有一个名字为 v，类型为 MethodPointersVertex 的接收者。

*/

package main

import (
	"fmt"
	"math"
)

type methodVertex struct {
	X, Y float64
}

func (v methodVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 记住：方法只是个带接收者参数的函数。
// 现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。
func Abs1(v methodVertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := methodVertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs1(v))
}
