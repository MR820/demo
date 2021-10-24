/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/10/13
 * Time 9:20 上午
 */

package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1
	//var (
	//	a = 1
	//	b = 1
	//)
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 4
	a, b = b, a
	t.Log(a)
	t.Log(b)
}
