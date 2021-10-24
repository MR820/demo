/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/8/25
 * Time 3:51 下午
 */

package main

import "fmt"

var a, b bool
var s string
var i int
var f float64

type A struct {
	a string
	b float32
}

func main() {
	x := A{
		a: "",
		b: 0,
	}
	fmt.Println(a, b, s, i, f, x)
}
