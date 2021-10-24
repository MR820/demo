/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/10/13
 * Time 上午9:57
 */

package main

import (
	"fmt"
)

func main() {
	Check(0)
}

func Check(n int) {

	defer func() {
		if recover() != nil {
			fmt.Println("is neither")
		}
	}()
	if Positive(n) {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

func Positive(n int) bool {
	if n == 0 {
		panic("undefined")
	}
	return n > -1
}
