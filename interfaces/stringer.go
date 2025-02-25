/**
* @File  : stringer.go
* @Author: Shide 2025/2/25 20:14
* @Desc  :
fmt 包中定义的 Stringer 是最普遍的接口之一。

type Stringer interface {
    String() string
}
Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。
*/

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{"张三", 23}
	b := Person{"李四", 23}
	fmt.Println(a, b)
}
func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
