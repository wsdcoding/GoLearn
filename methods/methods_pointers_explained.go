/**
@File  : methods_pointers_explained.go
@Author: Shide 2024/11/21 15:42
@Desc  : 指针与函数
参考methods_pointers.go 现在我们要把 Abs 和 Scale 方法重写为函数。

*/

package main

import (
	"fmt"
	"math"
)

type MethodsPoinntersExplainedVertex struct {
	X, Y float64
}

func Abs(v MethodsPoinntersExplainedVertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *MethodsPoinntersExplainedVertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := MethodsPoinntersExplainedVertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
