/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/11
 * Time 2:49 下午
 */

package main

import "fmt"

/**
677. 键值映射
实现一个 MapSum 类，支持两个方法，insert 和 sum：

MapSum() 初始化 MapSum 对象
void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。


示例：

输入：
["MapSum", "insert", "sum", "insert", "sum"]
[[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
输出：
[null, null, 3, null, 5]

解释：
MapSum mapSum = new MapSum();
mapSum.insert("apple", 3);
mapSum.sum("ap");           // return 3 (apple = 3)
mapSum.insert("app", 2);
mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)


提示：

1 <= key.length, prefix.length <= 50
key 和 prefix 仅由小写英文字母组成
1 <= val <= 1000
最多调用 50 次 insert 和 sum
*/
type MapSum struct {
	val  int
	next map[int32]*MapSum
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		val:  0,
		next: map[int32]*MapSum{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	node := this
	for _, v := range key {
		if _, ok := node.next[v]; !ok {
			trie := Constructor()
			node.next[v] = &trie
		}
		node = node.next[v]
	}
	node.val = val
}

func (this *MapSum) Sum(prefix string) int {
	//已前缀的最后一个节点为根节点，遍历所有节点求val和
	node := this
	for _, v := range prefix {
		if _, ok := node.next[v]; !ok {
			return 0
		}
		node = node.next[v]
	}
	//此时node为最后一个节点,以此为根。遍历
	val := 0
	s := make([]*MapSum, 0)
	s = append(s, node)
	for len(s) != 0 {
		a := s[0]
		s = s[1:]
		val += a.val
		for _, v := range a.next {
			s = append(s, v)
		}
	}
	return val
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

func main() {
	node := Constructor()
	root := &node
	root.Insert("apple", 3)
	fmt.Println(root.Sum("ap"))
	root.Insert("app", 2)
	fmt.Println(root.Sum("ap"))
}
