/**
@File  : defer_multi.go
@Author: Shide 2024/11/21 10:40
@Desc  : defer 栈
推迟调用的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的调用会按照后进先出的顺序调用。
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("计数：")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("结束")
}
