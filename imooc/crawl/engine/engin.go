/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/5/29
 * Time 4:40 下午
 */

package engine

import (
	"log"

	"imooc.com/ccmouse/learngo/imooc/crawl/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: err fetching url %s:%v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}
