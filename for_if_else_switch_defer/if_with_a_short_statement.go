/**
* @File  : if_with_a_short_statement.go
* @Author: Shide 2024/11/21 10:08
* @Desc  : if 和简短语句
 */

package main

import (
	"fmt"
	"math"
)

/*
*
和 for 一样，if 语句可以在条件表达式前执行一个简短语句。
该语句声明的变量作用域仅在 if 之内。
（在最后的 return 语句处使用 v 看看。）
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
