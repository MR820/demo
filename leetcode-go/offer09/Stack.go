/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/28
 * Time 1:11 下午
 */

package main

import "sync"

type node struct {
	value int
	prev  *node
}

type Stack struct {
	top    *node
	length int
	lock   *sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{
		top:    nil,
		length: 0,
		lock:   &sync.RWMutex{},
	}
}

func (this *Stack) Peek() int {
	if this.length == 0 {
		return -1
	}
	return this.top.value
}

func (this *Stack) Pop() int {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.length == 0 {
		return -1
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	n := &node{value, this.top}
	this.top = n
	this.length++
}
