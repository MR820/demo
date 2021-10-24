/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/26
 * Time 上午11:08
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"imooc.com/ccmouse/learngo/grpc/demo/data"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type demoServer struct {
	data.UnimplementedDemoServer
	savedResult []*data.Response // 用于服务端
}

var (
	tls  = flag.Bool("tls", false, "使用tls")
	port = flag.Int("port", 50050, "服务端口")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server.key")
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	s := grpc.NewServer(opts...)
	data.RegisterDemoServer(s, &demoServer{})

	reflection.Register(s)
	log.Printf("Server listening at :%v\n", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

// 实现方法add
func (s *demoServer) Add(ctx context.Context, in *data.TwoNum) (*data.Response, error) {
	x := in.X
	y := in.Y
	cRes := &data.Response{Result: x + y}
	return cRes, nil
}

// 实现方法GetStream
func (s *demoServer) GetStream(in *data.TwoNum, pipe data.Demo_GetStreamServer) error {
	err := pipe.Send(&data.Response{Result: in.X + in.Y})
	if err != nil {
		return err
	}
	err = pipe.Send(&data.Response{
		Result: in.X * in.Y,
	})
	if err != nil {
		return err
	}
	err = pipe.Send(&data.Response{Result: in.X - in.Y})
	if err != nil {
		return err
	}
	return nil
}

// PutStream
func (s *demoServer) PutStream(pipe data.Demo_PutStreamServer) error {
	var res int32
	for { // 循环接受
		request, err := pipe.Recv()
		if err == io.EOF { // 判断是否发送结束
			break
		}
		if err != nil {
			log.Println(err.Error())
		}
		res += request.X // 累加
	}
	_ = pipe.SendAndClose(&data.Response{Result: res}) // 返回
	return nil
}

// DoubleStream
func (s *demoServer) DoubleStream(pipe data.Demo_DoubleStreamServer) error {
	for {
		reques, err := pipe.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err = pipe.Send(&data.Response{Result: reques.X + reques.Y}); err != nil {
			return err
		}
	}
}
