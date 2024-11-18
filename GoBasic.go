package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//随机生成幸运数字
	fmt.Println("Like GO", rand.Intn(10))
	fmt.Printf("现在你有了 %g 个问题。\n", math.Sqrt(7))
}
