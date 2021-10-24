/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/1
 * Time 下午4:56
 */

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"imooc.com/ccmouse/learngo/grpc/tutorial"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	resp,err := http.Get("http://localhost:8080/protobuf")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		fmt.Println(err)
		return
	}
	p := &tutorial.Person{}
	proto.Unmarshal(body, p)
	fmt.Println(*p)
	fmt.Println(reflect.TypeOf(*p))
}
