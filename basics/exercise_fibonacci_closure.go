/**
@File  : exercise_fibonacci_closure.go
@Author: Shide 2024/11/21 15:25
@Desc  : 练习：斐波纳契闭包

*/

package main

import "fmt"

// fibonacci 函数返回一个闭包，该闭包返回斐波那契数列的下一个数
func fibonacci() func() int {
	var back1, back2 int = 0, 1 // 初始化前两个斐波那契数
	return func() int {
		var temp int = back1              // 保存当前的back1值
		back1, back2 = back2, back1+back2 // 更新back1和back2的值
		return temp                       // 返回当前的斐波那契数
	}
}

func main() {
	f := fibonacci()          // 调用fibonacci函数获取闭包
	for i := 0; i < 10; i++ { // 打印前10个斐波那契数
		fmt.Println(f())
	}
}
