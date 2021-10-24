/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/3/2
 * Time 2:16 下午
 */

package main

import "fmt"

type Human3 interface {
	Say(s string) error
	Eat(s string) error
	Walk(s string) error
}

type TestC string

func (t TestC) Say(s string) error {
	fmt.Printf("煎鱼：%s\n", s)
	return nil
}
func (t TestC) Eat(s string) error {
	fmt.Printf("煎鱼：%s\n", s)
	return nil
}

func (t TestC) Walk(s string) error {
	fmt.Printf("煎鱼：%s\n", s)
	return nil
}

func main() {
	var h Human3
	var t TestC
	h = t
	_ = h.Eat("烤羊排")
	_ = h.Say("炸鸡腿")
	_ = h.Walk("去炸鸡腿")
}
