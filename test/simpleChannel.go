/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/3/2
 * Time 11:41 上午
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan string)
	//go func() {
	//	ch <- "监狱风云"
	//}()
	//msg := <-ch
	//fmt.Println(msg)

	start := time.Now().Unix()
	ch := make(chan int)

	go func() {
		count := 0
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			count += i
		}
		ch <- count
	}()

	end := time.Now().Unix()
	fmt.Println(end - start)
	fmt.Println(<-ch)
	e := time.Now().Unix()
	fmt.Println(e - start)

}
