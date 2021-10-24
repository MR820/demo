/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/23
 * Time 下午5:50
 */

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	i := 10
	ip := &i
	fp := (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)
	fmt.Println(*fp)
	fmt.Println(reflect.TypeOf(fp))
}
