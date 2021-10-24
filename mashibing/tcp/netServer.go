/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/4
 * Time 10:43 上午
 */

package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8888")

	for {
		conn, _ := listener.Accept()

		fmt.Println("a client connected")
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		io.WriteString(conn, "hello")
		time.Sleep(1 * time.Second)
	}

}
