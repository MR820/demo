/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/23
 * Time 下午5:33
 */

package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var studentPool = sync.Pool{New: func() interface{} {
	return new(Student)
}}

func main() {
	var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})
	stu := studentPool.Get().(*Student)
	json.Unmarshal(buf, stu)
	studentPool.Put(stu)
	fmt.Println(stu)
}
