/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/16
 * Time 上午10:54
 */

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := []int{0, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Println(*pointer, " ")
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}
}
