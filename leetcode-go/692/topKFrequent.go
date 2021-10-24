/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/11
 * Time 4:24 下午
 */

package main

import (
	"fmt"
)

/**
692. 前K个高频单词
给一非空的单词列表，返回前 k 个出现次数最多的单词。

返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。

示例 1：

输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
输出: ["i", "love"]
解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
    注意，按字母顺序 "i" 在 "love" 之前。


示例 2：

输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
输出: ["the", "is", "sunny", "day"]
解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
    出现次数依次为 4, 3, 2 和 1 次。


注意：

假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。
输入的单词均由小写字母组成。


扩展练习：

尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。
*/

//大顶堆，弹出k个

type Heap struct {
	heap     []map[string]int
	limit    int
	heapSize int
}

func Constructor() *Heap {
	return &Heap{
		heap:     nil,
		limit:    0,
		heapSize: 0,
	}
}

func (this *Heap) isEmpty() bool {
	return this.heapSize == 0
}

func (this *Heap) isFull() bool {
	return this.heapSize == this.limit
}

func (this *Heap) push(m map[string]int) {
	if this.limit == this.heapSize {
		panic("heap is full")
	}
	this.heap[this.heapSize] = m
	heapInsert(this.heap, this.heapSize)
	this.heapSize += 1
}

func (this *Heap) pop() {

}

func heapInsert(m []map[string]int, index int) {

}

func topKFrequent(words []string, k int) []string {
	//首先遍历words，得到每个单词出现的频率
	//放入大小为k的大根堆中，弹出k次即可

}

func main() {
	s := []string{"plpaboutit", "jnoqzdute", "sfvkdqf", "mjc", "nkpllqzjzp", "foqqenbey", "ssnanizsav", "nkpllqzjzp",
		"sfvkdqf", "isnjmy", "pnqsz", "hhqpvvt", "fvvdtpnzx", "jkqonvenhx", "cyxwlef", "hhqpvvt", "fvvdtpnzx", "plpaboutit",
		"sfvkdqf", "mjc", "fvvdtpnzx", "bwumsj", "foqqenbey", "isnjmy", "nkpllqzjzp", "hhqpvvt", "foqqenbey", "fvvdtpnzx", "bwumsj",
		"hhqpvvt", "fvvdtpnzx", "jkqonvenhx", "jnoqzdute", "foqqenbey", "jnoqzdute", "foqqenbey", "hhqpvvt", "ssnanizsav", "mjc", "foqqenbey",
		"bwumsj", "ssnanizsav", "fvvdtpnzx", "nkpllqzjzp", "jkqonvenhx", "hhqpvvt", "mjc", "isnjmy", "bwumsj", "pnqsz", "hhqpvvt", "nkpllqzjzp",
		"jnoqzdute", "pnqsz", "nkpllqzjzp", "jnoqzdute", "foqqenbey", "nkpllqzjzp", "hhqpvvt", "fvvdtpnzx", "plpaboutit", "jnoqzdute", "sfvkdqf",
		"fvvdtpnzx", "jkqonvenhx", "jnoqzdute", "nkpllqzjzp", "jnoqzdute", "fvvdtpnzx", "jkqonvenhx", "hhqpvvt", "isnjmy", "jkqonvenhx", "ssnanizsav",
		"jnoqzdute", "jkqonvenhx", "fvvdtpnzx", "hhqpvvt", "bwumsj", "nkpllqzjzp", "bwumsj", "jkqonvenhx", "jnoqzdute", "pnqsz", "foqqenbey", "sfvkdqf",
		"sfvkdqf"}
	fmt.Println(topKFrequent(s, 1))
}
