/**
* @File  : type_assertions.go
* @Author: Shide 2025/2/25 20:08
* @Desc  :
 */

package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)

	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)

	f = i.(float64)
	fmt.Println(f)
}
