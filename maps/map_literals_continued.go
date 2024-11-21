/**
@File  : map_literals_continued.go
@Author: Shide 2024/11/21 15:04
@Desc  : 若顶层类型只是一个类型名，那么你可以在字面量的元素中省略它。

*/

package main

import "fmt"

type MapLiteralssVertex struct {
	Lat, Long float64
}

var maplitersrlsxx = map[string]MapLiteralssVertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(maplitersrlsxx)
}
