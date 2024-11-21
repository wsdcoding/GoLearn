/**
@File  : switch_with_no_condition.go
@Author: Shide 2024/11/21 10:37
@Desc  : 无条件 switch
无条件的 switch 同 switch true 一样。
这种形式能将一长串 if-then-else 写得更加清晰。
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}

}
