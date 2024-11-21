/**
* @File  : for_continued.go
* @Author: Shide 2024/11/21 9:05
* @Desc  : for 循环的初始化语句和后置语句是可选的。
 */

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
