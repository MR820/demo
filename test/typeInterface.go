/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/3/2
 * Time 2:08 下午
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var i interface{} = "吃煎鱼"
	//
	//s, ok := i.(string)
	//
	//if !ok {
	//	panic("error")
	//} else {
	//	fmt.Println(s)
	//}

	var i interface{} = "吃煎鱼"

	fmt.Println(reflect.TypeOf(i))
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String:
		fmt.Println("string")
	case reflect.Int:
		fmt.Println("string")
	default:
		fmt.Println("unhandled")

	}

	switch i.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("can not handler")

	}

	x := "煎鱼"
	var iface interface{} = x
	fmt.Println(iface)

}
