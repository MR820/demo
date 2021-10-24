/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/10/15
 * Time 下午3:19
 */

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var Counter int32 = 0

func main() {

	for routine := 1; routine <= 20; routine++ {
		wg.Add(1)
		go Routine(routine)
	}

	wg.Wait()
	fmt.Printf("Final Counter: %d \n", Counter)
}

func Routine(id int) {
	for count := 1; count <= 2; count++ {
		//value := Counter
		atomic.AddInt32(&Counter, 1)
		//Counter++
		//value++
		//Counter = value
	}
	wg.Done()
}
