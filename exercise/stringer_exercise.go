/**
* @File  : stringer_exercise.go
* @Author: Shide 2025/2/25 20:20
* @Desc  :

通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。

例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。
*/

package main

import "fmt"

type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr{
		"host": {192, 168, 1, 2},
		"dns":  {1, 1, 1, 1},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

}
