/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/9
 * Time 下午3:01
 */

package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	Read()
}

type Teacher struct {
	name  string
	age   int
	level string
}

type Student struct {
	name  string
	age   int
	class string
}

func (t *Teacher) Read() {
	fmt.Printf("teacher --- \n name:%v,age:%d,level:%v \n", t.name, t.age, t.level)
}

func (s *Student) Read() {
	fmt.Printf("student --- \n name:%v,age:%d,class:%v \n", s.name, s.age, s.class)
}

func main() {
	persons := map[string]interface{}{
		"teacher": Teacher{
			name:  "xzw",
			age:   30,
			level: "gj",
		},
		"student": Student{
			name:  "wz",
			age:   20,
			class: "2-2",
		},
	}

	for _, obj := range persons {
		fmt.Println(reflect.TypeOf(obj))
		t, isT := obj.(Teacher)
		if isT {
			t.Read()
		}
		s, isS := obj.(Student)
		if isS {
			s.Read()
		}
	}
}
