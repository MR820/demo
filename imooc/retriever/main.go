/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/12
 * Time 上午10:05
 */

package main

import (
	"fmt"

	"imooc.com/ccmouse/learngo/imooc/retriever/mock"
	"imooc.com/ccmouse/learngo/imooc/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this is a fake imooc.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
