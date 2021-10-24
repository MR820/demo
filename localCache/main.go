/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/1
 * Time 下午9:24
 */

package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"reflect"
	"time"
)

type MyStruct struct {
	name string
	age int
}

func main() {
	// 创建一个默认过期时间为 5 分钟的缓存，并且 每 10 分钟清除一次过期的项目
	c := cache.New(5*time.Minute, 10 * time.Minute)
	c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("baz", 42, cache.NoExpiration)
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
		fmt.Println(reflect.TypeOf(foo))
	}

	my := MyStruct{
		name: "xzw",
		age:  14,
	}

	c.Set("foo", my, cache.DefaultExpiration)
	if x, found := c.Get("foo"); found {
		fmt.Println(reflect.TypeOf(x))
		foo := x.(MyStruct)
		fmt.Println(foo)
	}
}