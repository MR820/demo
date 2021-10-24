/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/26
 * Time 下午4:22
 */

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type key string

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ctx = context.WithValue(ctx, key("name"), "hello")

	ch := make(chan context.Context, 0)

	go func() {
		go func() {
			defer func() {
				if v := recover(); v != nil {
					fmt.Printf("v:%v \n\r", v)
				}
			}()
			fmt.Println("B")

			cancel()

			fmt.Println(ctx.Value(key("name")))
			ctx = context.WithValue(ctx, key("user"), "hh")
			ch <- ctx

			time.Sleep(time.Second * 1)
			panic("error")

		}()
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("A")
		}
	}()
	fmt.Println("MAIN")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		fmt.Println("terminating: via signal")

	case <-ctx.Done():
		fmt.Println("error")

	case x := <-ch:
		fmt.Println(x.Value(key("user")))
		fmt.Println(x.Value(key("name")))
	}
}
