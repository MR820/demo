/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/30
 * Time 11:04 上午
 */

package main

import "fmt"

type maxHeap struct {
	heap     []int
	limit    int
	heapSize int
}

func Constructor(n int) *maxHeap {
	return &maxHeap{
		heap:     make([]int, n),
		limit:    n,
		heapSize: 0,
	}
}

func (h *maxHeap) isEmpty() bool {
	return h.heapSize == 0
}

func (h *maxHeap) isFull() bool {
	return h.limit == h.heapSize
}

func (h *maxHeap) push(v int) {
	if h.heapSize == h.limit {
		panic("heap is full")
	}
	h.heap[h.heapSize] = v
	heapInsert(h.heap, h.heapSize)
	h.heapSize += 1
}

func (h *maxHeap) pop() int {
	ans := h.heap[0]
	h.heapSize -= 1
	swap(h.heap, 0, h.heapSize)
	heapify(h.heap, 0, h.heapSize)
	return ans
}

func heapInsert(arr []int, idx int) {
	for arr[idx] > arr[(idx-1)/2] {
		swap(arr, idx, (idx-1)/2)
		idx = (idx - 1) / 2
	}
}

func heapify(arr []int, idx int, heapSize int) {
	left := idx*2 + 1
	largest := 0
	for left < heapSize {
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		} else {
			largest = left
		}
		if arr[largest] < arr[idx] {
			largest = idx
		}
		if largest == idx {
			break
		}
		swap(arr, largest, idx)
		idx = largest
		left = idx*2 + 1
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func findKthLargest(nums []int, k int) int {
	l := len(nums)
	h := Constructor(l)
	for _, v := range nums {
		h.push(v)
	}
	var ans int
	for i := 0; i < k; i++ {
		ans = h.pop()
	}
	return ans
}

func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(arr, 3))
}
