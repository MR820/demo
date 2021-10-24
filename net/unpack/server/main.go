/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/9
 * Time 下午5:34
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"imooc.com/ccmouse/learngo/net/unpack/znet"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("server listen err:", err)
		return
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		} else {
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("关闭")
	fmt.Println("新连接：", conn.RemoteAddr())

	dp := znet.NewDataPack()
	for {
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("read head err", err)
			break
		}
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}
		if msgHead.GetDataLen() > 0 {
			msg := msgHead
			msg.Data = make([]byte, msg.GetDataLen())

			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			fmt.Println("===> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ",data=", string(msg.Data))
		}
	}
}
