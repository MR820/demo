/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/1
 * Time 下午5:36
 */

package main

import (
	"context"
	"google.golang.org/grpc"
	"imooc.com/ccmouse/learngo/grpc/tutorial"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	tutorial.UnimplementedGreeterServer
}

func (s *server) GetPerson(ctx context.Context, in *tutorial.PersonId) (*tutorial.Person, error) {
	log.Printf("Received: %v", in.GetId())
	// 业务逻辑
	return &tutorial.Person{
			Name:        "",
			Id:          in.GetId(),
			Email:       "",
			Phones:      nil,
			LastUpdated: nil,
		},
		nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	ser := grpc.NewServer()
	tutorial.RegisterGreeterServer(ser, &server{})
	log.Printf("server listening as %v", lis.Addr())

	go func() {
		if err := ser.Serve(lis); err != nil {
			log.Fatalf("failed to server:%v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// 自动优雅停止，tcp避免占用资源
			ser.GracefulStop()
			return
		}
	}

}