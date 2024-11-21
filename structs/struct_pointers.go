/**
@File  : struct_pointers.go
@Author: Shide 2024/11/21 10:56
@Desc  : 结构体指针
结构体字段可通过结构体指针来访问。
如果我们有一个指向结构体的指针 p 那么可以通过 (*p).X 来访问其字段 X。
不过这么写太啰嗦了，所以语言也允许我们使用隐式解引用，直接写 p.X 就可以。

*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 //通常来说这是计算机中一种科学计数法的表示形式 1e9 = 110^9= 1000000000,例:9e8=910^8=900000000;e表示10，e后面的数字表示次方，e的多少次方。
	p.Y = 9e8
	fmt.Println(v)
}
