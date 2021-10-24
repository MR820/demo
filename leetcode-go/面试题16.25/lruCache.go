/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/16
 * Time 1:55 下午
 */

/**
面试题 16.25. LRU缓存
设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。

它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

示例:

LRUCache cache = new LRUCache( 2 );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
*/
package main

import "fmt"

type Node struct {
	k    int
	v    int
	pre  *Node
	next *Node
}

type LRUCache struct {
	head *Node
	tail *Node
	m    map[int]*Node
	cap  int
}

func Constructor(capacity int) LRUCache {
	head := &Node{
		v:    -1,
		pre:  nil,
		next: nil,
	}
	tail := &Node{
		v:    -1,
		pre:  nil,
		next: nil,
	}
	head.next = tail
	tail.pre = head

	return LRUCache{
		head: head,
		tail: tail,
		m:    map[int]*Node{},
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}
	//删除此node
	node := this.m[key]
	node.removeNode()
	this.addToHead(node)

	return node.v
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; ok {
		//存在则删除
		node := this.m[key]
		node.removeNode()
		//同时删除map
		delete(this.m, key)
	}
	if len(this.m) >= this.cap {

		//双联表删除尾
		rNode := this.removeFromTail()
		//map删除双联表尾巴的key
		delete(this.m, rNode.k)

	}

	this.m[key] = this.addToHead(&Node{
		k:    key,
		v:    value,
		pre:  nil,
		next: nil,
	})
}

func (this *LRUCache) removeFromTail() *Node {
	cur := this.tail.pre
	pre := cur.pre
	pre.next = this.tail
	this.tail.pre = pre
	return cur
}

func (this *LRUCache) addToHead(node *Node) *Node {

	head := this.head
	next := head.next

	head.next = node
	node.pre = head

	node.next = next
	next.pre = node
	return node
}

func (this *Node) removeNode() {
	node := this
	next := node.next
	pre := node.pre

	pre.next = next
	next.pre = pre

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))

	cache.Put(1, 3)
	fmt.Println(cache.Get(2))
	//
	//cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	//fmt.Println(cache.Get(3))
	//fmt.Println(cache.Get(4))

}
