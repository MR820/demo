/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/8
 * Time 上午10:35
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"imooc.com/ccmouse/learngo/net/protocol"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("err:", err)
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就推出
			return
		}
		data, err := protocol.Encode(inputInfo)
		if err != nil {
			panic(err)
		}
		_, err = conn.Write(data) // 发送数据
		if err != nil {
			return
		}
		reader := bufio.NewReader(conn)
		msg, err := protocol.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println(msg)
	}
}
