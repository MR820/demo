/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/17
 * Time 上午11:10
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
)

type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"` // 转发数
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

func main() {
	ctx := context.Background()

	client, err := elastic.NewClient(elastic.SetURL("http://node03:9200"))
	if err != nil {
		panic(err)
	}
	//
	info, code, err := client.Ping("http://node03:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion("http://node03:9200")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	exists, err := client.IndexExists("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		const mapping = `
		{
			"settings":{
				"number_of_shards":1,
				"number_of_replicas":0
			},
			"mappings": {
				"properties": {
					"user":{
						"type":"keyword"
					},
					"message":{
						"type":"text",
						"store":true,
						"fielddata":true
					},
					"retweets":{
						"type":"long"
					},
					"image":{
						"type":"keyword"
					},
					"created":{
						"type":"date"
					},
					"tag":{
						"type":"date"
					},
					"location":{
						"type":"geo_point"
					},
					"suggest_field":{
						"type":"completion"
					}
				}
			}
		}`
		createIndex, err := client.CreateIndex("twitter").
			BodyString(mapping).
			Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			// not acknowledged
		}
	}

	tweet1 := Tweet{
		User:     "olivere",
		Message:  "Take Five",
		Retweets: 0,
		Created:  time.Now(),
	}

	put1, err := client.Index().Index("twitter").Id("1").BodyJson(tweet1).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	tweet2 := `{"user":"olivere", "message":"It's a Raggy Waltz", "retweets":0, "created":"2021-08-01T14:43:05.840648+08:00"}`
	put2, err := client.Index().Index("twitter").Id("2").BodyString(tweet2).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	// Get tweet with specified ID
	get1, err := client.Get().
		Index("twitter").
		Id("1").
		Do(ctx)
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			panic(fmt.Sprintf("Document not found: %v", err))
		case elastic.IsTimeout(err):
			panic(fmt.Sprintf("Timeout retrieving document: %v", err))
		case elastic.IsConnErr(err):
			panic(fmt.Sprintf("Connection problem: %v", err))
		default:
			// Some other kind of error
			panic(err)
		}
	}
	msg1 := Tweet{}
	data, _ := get1.Source.MarshalJSON()
	json.Unmarshal(data, &msg1)
	fmt.Println(msg1)
	fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)

	// search with a term query
	searchResult, err := client.Search().
		Index("twitter").
		Query(elastic.NewTermQuery("user", "olivere")).
		Sort("user", true).
		From(0).
		Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		t := item.(Tweet)
		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	}
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	if searchResult.TotalHits() > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

		for _, hit := range searchResult.Hits.Hits {
			var t Tweet
			err := json.Unmarshal(hit.Source, &t)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	} else {
		fmt.Print("Found no tweets\n")
	}

	// update a tweet by the update API of Elasticsearch
	script := elastic.NewScript("ctx._source.retweets += params.num").Param("num", 1)
	update, err := client.Update().Index("twitter").Id("1").Script(script).
		Upsert(map[string]interface{}{"retweets": 0}).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New version of tweet %q is now %d", update.Id, update.Version)

	//// delete a tweet by id
	//_, err = client.Delete().
	//	Index("twitter").
	//	Id("1").
	//	Do(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// delete a tweet by termQuery
	//_, err = client.DeleteByQuery("twitter").
	//	Query(elastic.NewTermQuery("user", "olivere")).
	//	// 版本冲突继续执行
	//	ProceedOnVersionConflict().
	//	Do(ctx)
	//if err != nil {
	//	panic(err)
	//}

	////delete an index
	//deleteIndex, err := client.
	//	DeleteIndex("twitter").
	//	Do(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//if !deleteIndex.Acknowledged {
	//	// not acknowledged
	//	fmt.Println("not acknowledged")
	//}

	// 精确匹配单个字段
	searchResult, err = client.Search().
		Index("twitter").
		Query(elastic.NewTermQuery("message", "take")).
		Sort("created", true).
		From(0).
		Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var r Tweet
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(r)) {
			if t, ok := item.(Tweet); ok {
				fmt.Println(t)
			}
		}
	}

	// 通过terms实现SQL的in查询
	searchResult, err = client.Search().
		Index("twitter").
		Query(elastic.NewTermsQuery("message", "take", "five", "raggy")).
		From(0).
		Size(10).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var r Tweet
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(r)) {
			if t, ok := item.(Tweet); ok {
				fmt.Println(t)
			}
		}
	}

	// 模糊匹配单个字段
	searchResult, err = client.Search().
		Index("twitter").
		Query(elastic.NewMatchQuery("message", "take raggy")).
		From(0).
		Size(10).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var r Tweet
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(r)) {
			if t, ok := item.(Tweet); ok {
				fmt.Println(t)
			}
		}
	}

	// 范围查询
	searchResult, err = client.Search().
		Index("twitter").
		Query(elastic.NewRangeQuery("created").Gte("2021-08-01").Lt("2021-08-20")).
		From(0).
		Size(10).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var r Tweet
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(r)) {
			if t, ok := item.(Tweet); ok {
				fmt.Println(t)
			}
		}
	}

	// bool 组合查询
	// must
	// 创建 bool 查询条件
	boolQuery := elastic.NewBoolQuery().Must()
	// 创建 term 查询
	termQuery := elastic.NewTermQuery("user", "olivere")
	matchQuery := elastic.NewMatchQuery("message", "take a")
	// 设置bool查询的must条件, 组合了两个子查询
	// 表示搜索匹配user=olivere且message匹配"take a"的文档
	boolQuery.Must(termQuery, matchQuery)

	searchResult, err = client.Search().
		Index("twitter").
		Query(boolQuery).
		From(0).
		Size(10).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var r Tweet
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(r)) {
			if t, ok := item.(Tweet); ok {
				fmt.Println(t)
			}
		}
	}

	// must_not
	// 创建 bool 查询条件
	boolQuery = elastic.NewBoolQuery().Must()
	// 创建 term 查询
	//termQuery = elastic.NewTermQuery("user", "olivere")
	matchQuery = elastic.NewMatchQuery("message", "take five")

	boolQuery.MustNot(matchQuery)

	searchResult, err = client.Search().
		Index("twitter").
		Query(boolQuery).
		From(0).
		Size(10).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var r Tweet
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(r)) {
			if t, ok := item.(Tweet); ok {
				fmt.Println(t)
			}
		}
	}

	// should 条件
	boolQuery = elastic.NewBoolQuery().Must()
	// 创建 term 查询
	termQuery = elastic.NewTermQuery("user", "olivere")
	matchQuery = elastic.NewMatchQuery("message", "take five")

	boolQuery.Should(termQuery, matchQuery)

	searchResult, err = client.Search().
		Index("twitter").
		Query(boolQuery).
		From(0).
		Size(10).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var r Tweet
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(r)) {
			if t, ok := item.(Tweet); ok {
				fmt.Println(t)
			}
		}
	}
}
