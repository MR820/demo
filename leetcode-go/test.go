/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/13
 * Time 4:55 下午
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "abc"
	fmt.Println(reflect.TypeOf(s[0]))
	fmt.Println(s[0])
}
