/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/8
 * Time 上午11:36
 */

package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		log.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			log.Println("read udp failed, err:", err)
			continue
		}
		log.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			log.Println("write to udp failed, err:", err)
			continue
		}
	}
}
