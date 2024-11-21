/**
@File  : exercise_slices.go
@Author: Shide 2024/11/21 11:41
@Desc  : 练习：切片
实现 Pic。它应当返回一个长度为 dy 的切片，其中每个元素是一个长度为 dx，元素类型为 uint8 的切片。
当你运行此程序时，它会将每个整数解释为灰度值 （好吧，其实是蓝度值）并显示它所对应的图像。
图像的解析式由你来定。几个有趣的函数包括 (x+y)/2、x*y、x^y、x*log(y) 和 x%(y+1)。

（提示：需要使用循环来分配 [][]uint8 中的每个 []uint8。）

（请使用 uint8(intValue) 在类型之间转换；你可能会用到 math 包中的函数。）
*/

package main

import "math"
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	//创建二维切片
	result := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		// 为每个y创建一个长度为dx的切片
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8(x * int(math.Log(float64(y+1))))
		}
		result[y] = row
	}
	return result
}
func main() {
	pic.Show(Pic)
}
