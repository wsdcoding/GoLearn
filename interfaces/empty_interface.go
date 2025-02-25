/**
* @File  : empty_interface.go
* @Author: Shide 2025/2/25 20:03
* @Desc  :
 */

package main

import "fmt"

func main() {
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)
	i = "hello"
	describe2(i)
}
func describe2(i interface{}) {

	fmt.Printf("(%v, %T)\n", i, i)
}
