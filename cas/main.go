/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/26
 * Time 下午5:52
 */

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int32
	wg      sync.WaitGroup
)

func main() {

	threadNum := 1000
	wg.Add(threadNum)
	for i := 0; i < threadNum; i++ {
		go incr(i)
	}
	wg.Wait()
	fmt.Println(counter)
}

func incr(i int) {
	defer wg.Done()
	spinNum := 0
	for {
		old := counter
		ok := atomic.CompareAndSwapInt32(&counter, old, old+1)
		if ok {
			break
		} else {
			spinNum++
		}
	}
	if spinNum > 0 {
		fmt.Printf("thread,%d,spinNum,%d\n", i, spinNum)
	}

}
