/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/3/2
 * Time 1:46 下午
 */

package main

import "fmt"

type Human interface {
	Say(s string) error
}

type TestA string

func (t TestA) Say(s string) error {
	fmt.Printf("煎鱼:%s\n", s)
	return nil
}

func main() {
	var h Human
	var t TestA
	t.Say("烤鸡翅")
	h = t
	h.Say("烤羊排")
}
