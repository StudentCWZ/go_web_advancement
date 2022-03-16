/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/16 16:49
*/

package main

import (
	"fmt"
	"os"
)

// os.Args 操作
func main() {
	// os.Args 是一个 []string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}
