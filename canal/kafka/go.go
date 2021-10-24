/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/15
 * Time 下午1:43
 */

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	sig := make(chan struct{}, 1)

	for {
		sig <- struct{}{}
		go func() {

			kafkaReader()
			select {
			case <-sig:
				fmt.Println("consumer err")
			}
		}()
	}

	fmt.Println("kafka reader is end")
}

//type Pool struct {
//	m       sync.Mutex
//	res     chan *kafka.Reader
//	factory func() (*kafka.Reader, error)
//	closed  bool
//}

func kafkaReader() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:                []string{"node01:9092", "node02:9092", "node03:9092"},
		GroupID:                "test1",
		Topic:                  "example",
		QueueCapacity:          10000, // 内部消息队列的容量
		MinBytes:               10e3,
		MaxBytes:               10e6, // 10 M
		MaxWait:                10 * time.Second,
		ReadLagInterval:        0,
		GroupBalancers:         []kafka.GroupBalancer{kafka.RangeGroupBalancer{}},
		HeartbeatInterval:      3 * time.Second, // 心跳
		CommitInterval:         0,               // 提交offsets的间隔时间
		PartitionWatchInterval: 5 * time.Second, // 分区监控频率
		WatchPartitionChanges:  true,
		SessionTimeout:         30 * time.Second,       // 无心跳保持的时间长度
		RebalanceTimeout:       30 * time.Second,       // 成员加入协调器等待的时间长度
		JoinGroupBackoff:       5 * time.Second,        // 发生错误后的消费者组，重新加入等待时间
		RetentionTime:          24 * time.Hour,         // 消费者组被broker保存的时间长度
		StartOffset:            kafka.FirstOffset,      // FirstOffset or LastOffset
		ReadBackoffMin:         100 * time.Microsecond, // 在轮询新消息之前等待的最短时间
		ReadBackoffMax:         1 * time.Second,
		Logger:                 nil,
		ErrorLogger:            nil,
		IsolationLevel:         1, // ReadUncommitted IsolationLevel = 0 ReadCommitted   IsolationLevel = 1
		MaxAttempts:            3, // 消息传递错误之前将进行多少次尝试
	})
	defer r.Close()
	for {

		// 更细粒度的控制offset提交
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			panic(err)
		}

		// todo 此处为消息处理的逻辑
		log.Printf("topic:%s, partion:%d, offset:%d, value:%s, highwaterMark:%d",
			m.Topic, m.Partition, m.Offset, string(m.Value), m.HighWaterMark)

		if err := r.CommitMessages(context.Background(), m); err != nil {
			// 提交偏移量失败
			panic(err)
		}

		//
		//_, err := r.ReadMessage(context.Background())
		//if err != nil {
		//	panic(err)
		//}

	}
}
