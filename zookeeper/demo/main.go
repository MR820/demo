/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/23
 * Time 3:14 下午
 */

package main

import (
	"fmt"
	"time"

	"github.com/go-zookeeper/zk"
)

func callback(event zk.Event) {
	//fmt.Println(">>>>>>>>>>>>>>>>>>>")
	//fmt.Println("path:", event.Path)
	//fmt.Println("type:", event.Type.String())
	//fmt.Println("state:", event.State.String())
	//fmt.Println("<<<<<<<<<<<<<<<<<<<")
}

func main() {
	fmt.Printf("ZKOperateWatchTest\n")

	option := zk.WithEventCallback(callback)
	var hosts = []string{"172.16.213.131:2181", "172.16.213.132:2181", "172.16.213.133:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*10, option)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	var path1 = "/zk_test_go1"
	var data1 = []byte("zk_test_go1_data1")
	exist, s, _, err := conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("path[%s] exist[%t]\n", path1, exist)
	fmt.Printf("state:\n")
	fmt.Printf("%s\n", s)

	// try create
	var acls = zk.WorldACL(zk.PermAll)
	p, err_create := conn.Create(path1, data1, zk.FlagEphemeral, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}
	fmt.Printf("created path[%s]\n", p)

	time.Sleep(time.Second * 2)

	exist, s, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("path[%s] exist[%t] after create\n", path1, exist)
	fmt.Printf("state:\n")
	fmt.Printf("%s\n", s)

	// delete
	conn.Delete(path1, s.Version)

	exist, s, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("path[%s] exist[%t] after delete\n", path1, exist)
	fmt.Printf("state:\n")
	fmt.Printf("%s\n", s)
}
