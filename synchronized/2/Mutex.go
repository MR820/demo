/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/6/25
 * Time 下午3:42
 */

package main

import "sync"

var counter int

func incr(l *sync.Mutex) {
	l.Lock()
	for i := 0; i < 10; i++ {
		counter++
	}
	l.Unlock()
}
func main() {
	var l sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr(&l)
		}()
	}
	wg.Wait()
	println(counter)
}
