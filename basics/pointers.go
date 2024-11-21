/**
@File  : pointers.go
@Author: Shide 2024/11/21 10:43
@Desc  : 指针
Go 拥有指针。指针保存了值的内存地址。
类型 *T 是指向 T 类型值的指针，其零值为 nil。
var p *int
& 操作符会生成一个指向其操作数的指针。
i := 42
p = &i
* 操作符表示指针指向的底层值。
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
这也就是通常所说的「解引用」或「间接引用」。
与 C 不同，Go 没有指针运算。
*/

package main

import "fmt"

func main() {
	x := 10

	p := &x         //指针指向x
	fmt.Println(*p) //通过指针读取x的值
	*p = 11         // 通过指针赋值x的值
	fmt.Println(x)  //查看通过指针赋值后x的值
	fmt.Println(p)  //查看指针地址
	fmt.Println(*p) //查看指针所指的值

}
