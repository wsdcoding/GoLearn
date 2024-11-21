/**
@File  : struct_literals.go
@Author: Shide 2024/11/21 11:01
@Desc  : 结构体字面量
使用 Name: 语法可以仅列出部分字段（字段名的顺序无关）。

特殊的前缀 & 返回一个指向结构体的指针。

*/

package main

import "fmt"

type LITERALS struct {
	X, Y int
}

var (
	v1 = LITERALS{1, 3}  // 创建一个 MethodPointersVertex 类型的结构体
	v2 = LITERALS{X: 2}  // Y:0 被隐式地赋予零值
	v3 = LITERALS{}      // X:0 Y:0
	p  = &LITERALS{1, 5} // 创建一个 *MethodPointersVertex 类型的结构体（指针）

)

func main() {
	fmt.Println(v1, v2, v3, p)
	p.X = 6
	p.Y = 7
	fmt.Println(p.X, p.Y)
}
