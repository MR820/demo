/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/6/29
 * Time 上午9:59
 */

package main

import "fmt"

func strangeSlice() {
	//a := make([]int, 1, 10)
	//fmt.Printf("&a:%p,len:%d,cap:%d,a:%v\n", a, len(a), cap(a), a)
	//b := append(a, 1)
	//fmt.Printf("&b:%p,len:%d,cap:%d,b:%v\n", b, len(b), cap(b), b)
	//c := append(a, 2)
	//fmt.Printf("&c:%p,len:%d,cap:%d,c:%v\n", c, len(c), cap(c), c)
	//fmt.Printf("&b:%p,len:%d,cap:%d,b:%v\n", b, len(b), cap(b), b)
	//a = []int{1, 6, 9}
	//fmt.Printf("&a:%p,len:%d,cap:%d,a:%v\n", a, len(a), cap(a), a)
	// slice没有具体的初始化值
	a := []int{} // 或者var s []int
	// 后续使用append来增加元素
	a = append(a, 2, 4, 8, 9, 11, 12, 13, 15, 17)
	fmt.Printf("&a:%p,len:%d,cap:%d,a:%v\n", a, len(a), cap(a), a)
}

func praCopy() {
	s := []int{1, 2, 3}
	d := make([]int, 2, 5)
	// copy函数只会把目标slice中已初始化的元素修改为源slice中对应的元素值，而未初始化的元素不会被修改，依旧保持未初始化状态
	copy(d, s)
	// 当且仅当len(d) > len(s)时copy函数等价于以下代码
	/*
	   for i := range s {
	       d[i] = s[i]
	   }
	*/
	fmt.Printf("&d: %p, len: %d, cap: %d, d: %v\n", d, len(d), cap(d), d)
}

func strangeRange() {
	a := []int{1, 2, 3}
	// 以下代码遍历数组和遍历slice结果不一样
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 200, 300
			fmt.Println(a)
		}
		a[i] = v + 100
	}
	fmt.Println(a)
}

func praAssignment() {
	// 可将a改为slice看看运行结果
	a := [3]int{1, 2, 3}
	b := a // a为数组时会发生内存拷贝, a和b的内容相同，但是内存地址不同!!!
	fmt.Printf("[before]&a: %p,a: %v; &b: %p, b: %v\n", &a[0], a, &b[0], b)
	b[0] = 6
	fmt.Printf("[after]&a: %p,a: %v; &b: %p, b: %v\n", &a, a, &b, b)

}

func main() {
	strangeSlice()
	praCopy()
	strangeRange()
	praAssignment()
}
