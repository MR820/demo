/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/3/2
 * Time 2:03 下午
 */

package main

import "fmt"

type Human2 interface {
	Say(s string) error
	Eat(s string) error
}

type TestB struct {
}

func (t TestB) Say(s string) error {
	fmt.Printf("说煎鱼:%s\n", s)
	return nil
}

func (t *TestB) Eat(s string) error {
	fmt.Printf("吃煎鱼:%s\n", s)
	return nil
}

func main() {
	var h Human2 = &TestB{}
	h.Say("催更")
	h.Eat("真香")
}
