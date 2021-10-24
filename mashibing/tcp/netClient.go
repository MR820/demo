/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/4
 * Time 10:40 上午
 */

package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8888")
	//fmt.Println(conn)
	//log.Println("connected")
	defer conn.Close()
	io.Copy(os.Stdout, conn)
}
