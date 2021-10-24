/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/3/2
 * Time 11:07 上午
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	rv := []interface{}{"hi", 42, func() {}}
	for _, v := range rv {
		//fmt.Println(reflect.TypeOf(v))
		//fmt.Println(reflect.ValueOf(v).Kind())
		//fmt.Println(reflect.Int)

		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())

		}
	}
}
