/**
@File  : struct.go
@Author: Shide 2024/11/21 10:49
@Desc  : 结构体
一个 结构体（struct）就是一组 字段（field）。
*/

package main

import "fmt"

type Shide struct {
	name string
	age  int
}

func main() {
	fmt.Println(Shide{"shidea", 27})
}
