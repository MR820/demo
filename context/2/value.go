/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/16
 * Time 下午3:14
 */

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	go helloHandle(ctx, 2*time.Second)
	go f2(ctx)
	select {
	case <-ctx.Done(): //此处阻塞2s
		fmt.Println("Hello Handle", ctx.Err())
	}
	time.Sleep(3 * time.Second)
}

func helloHandle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.Tick(duration):
		fmt.Println("process request with", duration)
	}
}

func f2(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("f2", ctx.Err())
	}
}
