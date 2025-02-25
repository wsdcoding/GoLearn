/**
* @File  : nil_interface_value.go
* @Author: Shide 2025/2/25 19:55
* @Desc  :
nil 接口值既不保存值也不保存具体类型。

为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。
*/

package main

import "fmt"

type Ik interface {
	M()
}

func describle(i Ik) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func main() {
	var i Ik
	describle(i)
	i.M()
}
