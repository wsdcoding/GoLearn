/**
* @File  : errors.go
* @Author: Shide 2025/2/25 20:27
* @Desc  :
Go 程序使用 error 值来表示错误状态。

与 fmt.Stringer 类似，error 类型是一个内建接口：

type error interface {
    Error() string
}
（与 fmt.Stringer 类似，fmt 包也会根据对 error 的实现来打印值。）

通常函数会返回一个 error 值，调用它的代码应当判断这个错误是否等于 nil 来进行错误处理。

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
error 为 nil 时表示成功；非 nil 的 error 表示失败。
*/

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
