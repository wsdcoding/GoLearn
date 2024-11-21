package main

import "fmt"

/*
*
变量声明可以包含初始值，每个变量对应一个。

如果提供了初始值，则类型可以省略；变量会从初始值中推断出类型。
*/
var i, j int = 1, 2

func main() {
	var c, py, java = true, false, "no!"
	fmt.Println(i, j, c, py, java)
}
