/**
* @File  : interface_values_with_nil.go
* @Author: Shide 2025/2/25 19:46
* @Desc  :
即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。

注意: 保存了 nil 具体值的接口其自身并不为 nil。
*/

package main

import "fmt"

type IM interface {
	M()
}
type TM struct {
	S string
}

func (t *TM) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i IM) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func main() {
	var i IM
	var t *TM
	i = t
	describe(i)
	i.M()

	i = &TM{"你好"}
	describe(i)
	i.M()
}
