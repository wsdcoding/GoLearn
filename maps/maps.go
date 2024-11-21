/**
@File  : maps.go
@Author: Shide 2024/11/21 14:54
@Desc  : Map映射
map 映射将键映射到值。
映射的零值为 nil 。nil 映射既没有键，也不能添加键。
make 函数会返回给定类型的映射，并将其初始化备用。
*/

package main

import "fmt"

type MapVertex struct {
	Lat, Long float64
}

var m map[string]MapVertex

func main() {
	m = make(map[string]MapVertex)
	m["Bell Labs"] = MapVertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
