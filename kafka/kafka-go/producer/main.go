/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/20
 * Time 上午9:34
 */

package main

import (
	"context"
	"fmt"
	"learngo/kafka/logd"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	var wg = sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			kafkaWriter()
		}()
	}
	wg.Wait()
	fmt.Println("生产消息成功")

}

func callback(messages []kafka.Message, err error) {
	if err != nil {
		for _, v := range messages {
			s := fmt.Sprintf(string(v.Value))
			logd.WriteLog("error.log", s)
		}
		panic(err)
	} else {
		for _, v := range messages {
			s := fmt.Sprintf("partion : %d, offset : %d, value : %v \n", v.Partition, v.Offset, string(v.Value))
			logd.WriteLog("msg.log", s)
		}
	}
}

func kafkaWriter() {

	w := kafka.Writer{
		Addr:         kafka.TCP("node01:9092", "node02:9092", "node03:9092"),
		Topic:        "test",
		Balancer:     &kafka.LeastBytes{},
		MaxAttempts:  3, // 限制发送消息的尝试次数
		BatchSize:    10000,
		BatchBytes:   10485760, // 10M
		BatchTimeout: 1 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		RequiredAcks: -1,
		Async:        true, // 异步，调用者不会收到返回值。仅在不关心消息是否写入kafka的场景下使用
		Completion:   callback,
		Compression:  0,
		Logger:       nil,
		ErrorLogger:  nil,
		Transport:    nil,
	}
	defer w.Close()

	//key := fmt.Sprintf("key-%d", num)

	var msg string
	for i := 0; i < 1000; i++ {
		msg = fmt.Sprintf("test %d", i)
		err := w.WriteMessages(
			context.Background(),
			kafka.Message{
				//Topic:     "test",
				//Partition: 0,
				//Offset:        0,
				//HighWaterMark: 0,
				Key:   []byte(string(i)),
				Value: []byte(msg),
				//Headers: nil,
				//Time:    time.Time{},
			},
		)
		if err != nil {
			panic(err)
		}
	}
}
