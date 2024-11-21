/**
@File  : struct_fields.go
@Author: Shide 2024/11/21 10:52
@Desc  : 结构体字段

*/

package main

import "fmt"

type Shide1 struct {
	name string
	age  int
}

func main() {
	x := Shide1{"shideaaaaaa", 28}
	fmt.Println(x.name)
	x.name = "shide调用重写 "
	fmt.Println(x.name)
}
