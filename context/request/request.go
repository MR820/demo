/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/9
 * Time 下午5:26
 */

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 100*time.Millisecond)
	req, _ := http.NewRequest(http.MethodGet, "http://www.google.com", nil)
	req = req.WithContext(ctx)
	client := &http.Client{}
	res, err := client.Do(req)

	//res, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	fmt.Println("response received, status code:", res.StatusCode)
}
