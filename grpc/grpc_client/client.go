/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/1
 * Time 下午5:47
 */

package main

import (
	"context"
	"google.golang.org/grpc"
	"imooc.com/ccmouse/learngo/grpc/tutorial"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	address = "localhost:8000"
	defaultId = 123
)


func main() {
	conn,err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := tutorial.NewGreeterClient(conn)

	var id int32 = defaultId
	if len(os.Args) > 1  {
		x, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err != nil {
			log.Fatalf("param illgal: %v", err)
		}
		id = int32(x)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPerson(ctx, &tutorial.PersonId{Id:id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %d", r.GetId())
}