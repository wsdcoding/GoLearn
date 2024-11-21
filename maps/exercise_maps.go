/**
@File  : exercise_maps.go
@Author: Shide 2024/11/21 15:09
@Desc  : 练习：映射
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。 函数 wc.Test 会为此函数执行一系列测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有用。

*/

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	mmap := make(map[string]int)
	arraymap := strings.Fields(s)
	for _, v := range arraymap {
		mmap[v] = mmap[v] + 1
	}
	return mmap
}
func main() {
	wc.Test(WordCount)
}
