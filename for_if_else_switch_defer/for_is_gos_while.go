/**
* @File  : for_is_gos_while.go
* @Author: Shide 2024/11/21 9:08
* @Desc  : for 是 Go 中的 while
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
