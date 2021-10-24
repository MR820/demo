/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/16
 * Time 上午10:53
 */

package main

import (
	"fmt"
	"unsafe"
)

type Teacher struct {
	name string
	age  int
}

func main() {
	t := Teacher{"ttt", 20}

	pt := unsafe.Pointer(&t)
	name := (*string)(unsafe.Pointer(pt))
	fmt.Println("name:", *name)
}
