/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/16
 * Time 下午5:40
 */

package main

import (
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {

	hystrix.ConfigureCommand("get_baidu", hystrix.CommandConfig{
		Timeout:               100,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 1,
	})
	output := make(chan bool, 0)
	hystrix.Go("get_baidu", func() error {
		_, err := http.Get("https://www.baidu.com/")

		output <- true
		if err != nil {
			fmt.Println("get error")
			return err
		}
		return nil
	}, func(err error) error {
		fmt.Println("get an error, handle it", err)
		output <- false
		return nil
	})

	select {
	case out := <-output:
		fmt.Println(out)
	}

}
