/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/6
 * Time 上午11:31
 */

package main

import (
	"fmt"
	"reflect"
)

type Ref struct {
	id   int
	name string
}

func (ref *Ref) GetName() {
	fmt.Println("getName()函数")
}
func (ref *Ref) GetNameById() {
	fmt.Println("getNameById()函数")
}
func main() {
	ref := new(Ref)
	t := reflect.TypeOf(ref)
	v := reflect.ValueOf(ref)
	fmt.Println(t)
	fmt.Println(v)
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
		v.Method(i).Call(nil)
	}
}
