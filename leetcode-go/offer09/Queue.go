/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/28
 * Time 1:10 下午
 */

/**
剑指 Offer 09. 用两个栈实现队列
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )



示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
*/

package main

import "fmt"

type CQueue struct {
	in  *Stack
	out *Stack
}

func Constructor() CQueue {
	return CQueue{
		in:  NewStack(),
		out: NewStack(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.out.length > 0 {
		return this.out.Pop()
	}
	var v int
	for this.in.length > 0 {
		v = this.in.Pop()
		this.out.Push(v)
	}
	return this.out.Pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	c := Constructor()
	q := &c
	fmt.Println(q.DeleteHead())
	q.AppendTail(5)
	q.AppendTail(2)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
}
