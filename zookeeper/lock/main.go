/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/15
 * Time 下午2:52
 */

package main

import (
	"time"

	"github.com/go-zookeeper/zk"
)

func main() {
	var servers = []string{"172.16.213.131:2181", "172.16.213.132:2181", "172.16.213.133:2181"}
	c, _, err := zk.Connect(servers, time.Second*10)
	if err != nil {
		panic(err)
	}
	//fmt.Println(zk.PermRead)
	//fmt.Println(zk.PermWrite)
	//fmt.Println(zk.PermCreate)
	//fmt.Println(zk.PermDelete)
	//fmt.Println(zk.PermAdmin)
	//fmt.Println(zk.PermAll)

	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err = l.Lock()
	if err != nil {
		panic(err)
	}
	println("lock success")

	time.Sleep(time.Second * 10)
	l.Unlock()
}
