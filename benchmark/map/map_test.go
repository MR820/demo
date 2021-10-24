/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/23
 * Time 下午5:16
 */

package _map

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func timeGC() time.Duration {
	start := time.Now()
	runtime.GC()
	return time.Since(start)
}

func mapPointer(num int) {
	m := make(map[int]*int, num)
	for i := 0; i < num; i++ {
		m[i] = &i
	}
	runtime.GC()
	fmt.Printf("With point %T, GC took %s\n", m, timeGC())
	_ = m[0]
}

func mapValue(num int) {
	m := make(map[int]int, num)
	for i := 0; i < num; i++ {
		m[i] = i
	}
	runtime.GC()
	fmt.Printf("With %T, GC took %s\n", m, timeGC())
	_ = m[0]
}

func mapPointerShard(num int) {
	shards := make([]map[int]*int, 100)
	for i := range shards {
		shards[i] = make(map[int]*int)
	}
	for i := 0; i < num; i++ {
		shards[i%100][i] = &i
	}
	runtime.GC()
	fmt.Printf("With map shards (%T), GC took %s\n", shards, timeGC())
	_ = shards[0][0]
}

func mapValueShard(num int) {
	shards := make([]map[int]int, 100)
	for i := range shards {
		shards[i] = make(map[int]int)
	}
	for i := 0; i < num; i++ {
		shards[i%100][i] = i
	}
	runtime.GC()
	fmt.Printf("With map shards (%T), GC took %s\n", shards, timeGC())
	_ = shards[0][0]
}

const N = 5e7 // 5000w

func BenchmarkMapPointer(b *testing.B) {
	mapPointer(N)
}

func BenchmarkMapValue(b *testing.B) {
	mapValue(N)
}

func BenchmarkMapPointerShard(b *testing.B) {
	mapPointerShard(N)
}

func BenchmarkMapValueShard(b *testing.B) {
	mapValueShard(N)
}
