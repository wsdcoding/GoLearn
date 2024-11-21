/**
* @File  : for.go
* @Author: Shide 2024/11/21 9:03
* @Desc  : for 循环
 */

package main

import "fmt"

/*
*
Go 只有一种循环结构：for 循环。
基本的 for 循环由三部分组成，它们用分号隔开：
初始化语句：在第一次迭代前执行
条件表达式：在每次迭代前求值
后置语句：在每次迭代的结尾执行
初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。

一旦条件表达式求值为 false，循环迭代就会终止。
*/
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
