/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/16
 * Time 上午10:51
 */

package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	s := Student{}
	s.Name = "Peter"
	s.Age = 33

	pStudent := unsafe.Pointer(&s)
	// 整个对象转换成指针，默认是获取第一个属性
	name := (*string)(unsafe.Pointer(pStudent))
	fmt.Println("name:", *name)
	// 利用Offsetof获取age属性的偏移量获取属性
	age := (*int)(unsafe.Pointer(uintptr(pStudent) + unsafe.Offsetof(s.Age)))
	fmt.Println("age:", *age)

	// 修改指针的值
	*name = "Mary"
	*age = 20
	fmt.Println(s)
}
