/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/9
 * Time 下午6:03
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go f1(ch)
	go f2(ch)
	go f3(ch)
	ch <- 1
	time.Sleep(1 * time.Second)
	close(ch)

	time.Sleep(5 * time.Second)
	r, ok := <-ch
	fmt.Println(ok)
	fmt.Println(r)
}

func f1(ch chan int) {
	select {
	case r := <-ch:
		fmt.Println(r)
	}
}
func f2(ch chan int) {
	select {
	case r := <-ch:
		fmt.Println(r)
	}
}
func f3(ch chan int) {
	select {
	case r := <-ch:
		fmt.Println(r)
	}
}
