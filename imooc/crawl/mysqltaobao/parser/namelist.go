/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/5/29
 * Time 4:33 下午
 */

package parser

import (
	"regexp"

	"imooc.com/ccmouse/learngo/imooc/crawl/engine"
)

var NameListRe string = `<a target="_top" class="main" href="(/monthly/[^"]+)"[^>]*>([^<]+)</a>`
var prefix string = "http://mysql.taobao.org"

func ParseNameList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(NameListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        prefix + string(m[1]),
			ParserFunc: ParseList,
		})
	}
	return result
}
