/**
@File  : map_literals.go
@Author: Shide 2024/11/21 14:59
@Desc  : 映射字面量
映射的字面量和结构体类似，只不过必须有键名。

*/

package main

import "fmt"

type MapLiteralVertex struct {
	Lat, Long float64
}

var map_liters = map[string]MapLiteralVertex{
	"Bell Labs": MapLiteralVertex{
		40.68433, -74.39967,
	},
	"Google": MapLiteralVertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(map_liters)
}
