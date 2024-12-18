/**
* @File  : if.go
* @Author: Shide 2024/11/21 9:10
* @Desc  : if 判断，Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( )，而大括号 { } 则是必须的
 */

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

}
