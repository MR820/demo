/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/15
 * Time 下午3:26
 */

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	address    = "172.16.213.134:6379"
	expireTime = 600
)

type RemoteSpikeKeys struct {
	SpikeOrderHashKey  string
	TotalInventoryKey  string
	QuantityOfOrderKey string
}

const LuaScript = `
		local product_key = KEYS[1]
		local product_total_key = ARGV[1]
		local product_sold_key = ARGV[2]
		local product_total_nums = tonumber(redis.call('HGET', product_key, product_total_key))
		local product_sold_nums = tonumber(redis.call('HGET', product_key, product_sold_key))
		if (product_total_nums > product_sold_nums) then
			return redis.call('HINCRBY', product_key, product_sold_key, 1)
		end
		return 0
	`

func main() {
	r := RemoteSpikeKeys{
		SpikeOrderHashKey:  "sp001",
		TotalInventoryKey:  "total",
		QuantityOfOrderKey: "sold",
	}

	redisClient := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow:    nil,
		MaxIdle:         100,
		MaxActive:       1000,
		IdleTimeout:     600 * time.Second,
		Wait:            true,
		MaxConnLifetime: 600 * time.Second,
	}

	time.Sleep(5 * time.Second)

	wg := sync.WaitGroup{}

	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			rc := redisClient.Get()
			defer rc.Close()
			r.RemoteDeductionStock(rc)
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func (r *RemoteSpikeKeys) RemoteDeductionStock(conn redis.Conn) bool {
	lua := redis.NewScript(1, LuaScript)
	result, err := redis.Int(lua.Do(conn, r.SpikeOrderHashKey, r.TotalInventoryKey, r.QuantityOfOrderKey))
	if err != nil {
		panic(err)
	}
	return result != 0
}
