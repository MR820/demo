/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/11
 * Time 10:26 上午
 */

/**
208. 实现 Trie (前缀树)
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
说明:

你可以假设所有的输入都是由小写字母 a-z 构成的。
保证所有输入均为非空字符串。

*/

package main

import (
	"fmt"
)

type Trie struct {
	end  int
	next [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		end:  0,
		next: [26]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		index := v - 'a'
		if node.next[index] == nil {
			new := Constructor()
			node.next[index] = &new
		}
		node = node.next[index]
	}
	node.end++
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, v := range word {
		index := v - 'a'
		if node.next[index] == nil {
			return false
		}
		node = node.next[index]
	}
	if node.end > 0 {
		return true
	} else {
		return false
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix {
		index := v - 'a'
		if node.next[index] == nil {
			return false
		}
		node = node.next[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	new := Constructor()
	root := &new
	root.Insert("apple")
	fmt.Println(root.Search("apple"))
	fmt.Println(root.Search("app"))
	fmt.Println(root.StartsWith("app"))
	root.Insert("app")
	fmt.Println(root.Search("app"))

}
