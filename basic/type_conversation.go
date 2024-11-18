/**
* @File  : type_conversation.go
* @Author: Shide 2024/11/18 17:09
* @Desc  : 类型转换
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 1, 2
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
