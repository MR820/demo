/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/9
 * Time 下午5:46
 */

package main

import (
	"fmt"
	"net"
	"time"

	"imooc.com/ccmouse/learngo/net/unpack/znet"
)

func main() {

	conn, err := net.DialTimeout("tcp", "localhost:8080", time.Second*5)
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}
	dp := znet.NewDataPack()
	msg1 := &znet.Message{
		Id:      0,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err:", err)
		return
	}
	msg2 := &znet.Message{
		Id:      1,
		DataLen: 7,
		Data:    []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)
	select {}
}
