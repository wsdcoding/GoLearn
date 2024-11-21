/**
@File  : defer.go
@Author: Shide 2024/11/21 10:38
@Desc  : defer 推迟
defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
*/

package main

import "fmt"

func main() {

	fmt.Println("hello defer")
	defer fmt.Println("defer")
	fmt.Println("hello")

}
