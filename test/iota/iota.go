/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/7/8
 * Time 下午6:09
 */

package main

import "fmt"

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

func main() {
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
	//fmt.Println(mutexWaiterShift)
}
