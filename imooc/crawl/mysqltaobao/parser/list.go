/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/5/29
 * Time 5:46 下午
 */

package parser

import (
	"regexp"

	"imooc.com/ccmouse/learngo/imooc/crawl/engine"
)

var ListRe string = `<a href="(/monthly/[^"]+)" target="_blank">([^<]+)</a>`

func ParseList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(ListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        prefix + string(m[1]),
			ParserFunc: ParseContent,
		})
	}
	return result
}
