/**
@File  : methods_with_pointer_receivers.go
@Author: Shide 2024/11/21 16:12
@Desc  : 选择值或指针作为接收者
使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样会更加高效。

在本例中，Scale 和 Abs 接收者的类型为 *MWPRVertex，即便 Abs 并不需要修改其接收者。

通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。 （我们会在接下来几页中明白为什么。）

*/

package main

import (
	"fmt"
	"math"
)

type MWPRVertex struct {
	X, Y float64
}

func (v *MWPRVertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *MWPRVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &MWPRVertex{3, 4}
	fmt.Printf("缩放前：%+v，绝对值：%v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("缩放后：%+v，绝对值：%v\n", v, v.Abs())
}
