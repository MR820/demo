/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/5/29
 * Time 5:07 下午
 */

package parser

import (
	"fmt"
	"regexp"

	"imooc.com/ccmouse/learngo/imooc/crawl/engine"
)

var ContentRe string = `<div class="title"><h2>(.+)</h2></div>`

func ParseContent(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(ContentRe)
	matches := re.FindAllSubmatch(contents, -1)

	for _, ma := range matches {
		fmt.Printf("%s", string(ma[0]))
		fmt.Printf("%s", string(ma[1]))
	}

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        prefix + string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
