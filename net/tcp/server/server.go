/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/8
 * Time 上午10:27
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	"imooc.com/ccmouse/learngo/net/protocol"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080") // 监听端口
	if err != nil {
		log.Fatal("listen failed, err:", err)
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			log.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个 goroutine 处理连接
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)

		msg, err := protocol.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
		data, err := protocol.Encode(msg)
		if err != nil {
			log.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
