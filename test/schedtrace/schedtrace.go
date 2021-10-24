/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/5
 * Time 下午5:26
 */

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int
var l sync.Mutex

func incr() {
	l.Lock()
	counter++
	l.Unlock()
}

func printMemStats() {
	t := time.NewTicker(time.Millisecond)
	s := runtime.MemStats{}
	for {
		select {
		case <-t.C:
			runtime.ReadMemStats(&s)
			fmt.Printf("gc %d last@%v, next_heap_size@%vMB\n", s.NumGC,
				time.Unix(int64(time.Duration(s.LastGC).Seconds()), 0), s.NextGC/(1<<20))
		}
	}
}

func main() {
	//f, err := os.Create("trace.out")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//
	//err = trace.Start(f)
	//if err != nil {
	//	panic(err)
	//}
	//defer trace.Stop()

	go printMemStats()
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
	println(counter)

}
