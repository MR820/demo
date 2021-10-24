/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/3/2
 * Time 2:37 下午
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now().Unix())
	//timer := time.NewTimer(2 * time.Second)
	//<-timer.C
	//fmt.Println("煎鱼")

	/**
	v := make(chan struct{})
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("我想吃煎鱼")
		v <- struct{}{}
	})
	defer timer.Stop()
	<-v

	*/

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done")
			return
		case t := <-ticker.C:
			fmt.Println("炸煎鱼", t.Unix())

		}
	}
}
