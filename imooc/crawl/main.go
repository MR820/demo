/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/5/29
 * Time 3:38 下午
 */

package main

import (
	"imooc.com/ccmouse/learngo/imooc/crawl/engine"
	"imooc.com/ccmouse/learngo/imooc/crawl/mysqltaobao/parser"
)

var url string = "http://mysql.taobao.org/monthly/2014/08/01/"

func main() {
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseContent,
	})

}
