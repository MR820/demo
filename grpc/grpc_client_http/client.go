/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/1
 * Time 下午7:11
 */

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"imooc.com/ccmouse/learngo/grpc/tutorial"
	"log"
	"net/http"
	"time"
)

type Server struct {
	engine *gin.Engine
	webClient tutorial.GreeterClient
}

func InitWebServer (addr string) tutorial.GreeterClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn,err := grpc.DialContext(ctx, addr, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := tutorial.NewGreeterClient(conn)
	return client
}

func (s *Server) CallWeb(txt *gin.Context) {
	var in *tutorial.PersonId
	if err := txt.ShouldBind(&in); err != nil {
		log.Fatalf(err.Error())
	}
	out, err := s.webClient.GetPerson(context.Background(), in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	txt.JSON(http.StatusOK, out)
}

func NewHttp(httpAddr string, rpcAddr string) {
	r := gin.Default()
	s := Server{
		engine:    r,
		webClient: InitWebServer(rpcAddr),
	}
	s.engine.GET("/hello", s.CallWeb)
	r.Run(httpAddr)
}

func main() {
	NewHttp(":8080", ":8000")
}