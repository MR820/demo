/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/3/2
 * Time 11:35 上午
 */

package main

import (
	"log"
	"reflect"
)

func main() {
	i := 2.33
	v := reflect.ValueOf(&i)
	v.Elem().SetFloat(3.14159)
	log.Println("value:", i)
}
