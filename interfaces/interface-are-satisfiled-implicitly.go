/**
* @File  : interface-are-satisfiled-implicitly.go
* @Author: Shide 2025/1/18 19:43
* @Desc  : 接口与隐式实现
类型通过实现一个解耦的所有方法来实现该接口，既然无需专门显示生命，也就没有implement关键字
隐式接口从接口的视线中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备，因此，也就无需再每一个实现上增加新的接口名称，这样同事也鼓励了明确的接口定义。
*/

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}
func main() {
	var i I = T{"hello"}
	i.M()
}
