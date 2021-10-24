/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/18
 * Time 5:50 下午
 */
/**
225. 用队列实现栈
使用队列实现栈的下列操作：

push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
*/
package main

import "fmt"

type MyStack struct {
	a, b MyQueue
}

func Constructor() MyStack {
	return MyStack{
		a: NewMyQueue(),
		b: NewMyQueue(),
	}
}

func (this *MyStack) Push(data int) {
	this.a.Enqueue(data)
	for this.b.size() > 0 {
		this.a.Enqueue(this.b.Dequeue())
	}
	this.a, this.b = this.b, this.a
}

func (this *MyStack) Pop() int {
	return this.b.Dequeue()
}

func (this *MyStack) Top() int {
	res := this.Pop()
	this.Push(res)
	return res
}

func (this *MyStack) Empty() bool {
	return this.b.isEmpty()
}

type MyQueue struct {
	data []int
}

func NewMyQueue() MyQueue {
	return MyQueue{data: nil}
}

func (this *MyQueue) Enqueue(data int) {
	this.data = append(this.data, data)
}

func (this *MyQueue) Dequeue() int {
	res := this.data[0]
	if !this.isEmpty() {
		this.data = this.data[1:]
	}
	return res

}

func (this *MyQueue) isEmpty() bool {
	if len(this.data) == 0 {
		return true
	} else {
		return false
	}
}

func (this *MyQueue) size() int {
	return len(this.data)
}

func main() {
	head := Constructor()
	head.Push(1)
	head.Push(2)
	fmt.Println(head.Top())
	fmt.Println(head.Pop())
	fmt.Println(head.Empty())

}
