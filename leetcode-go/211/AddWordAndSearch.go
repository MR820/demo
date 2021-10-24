/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/10
 * Time 4:01 下午
 */

/**
211. 添加与搜索单词 - 数据结构设计
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

实现词典类 WordDictionary ：

WordDictionary() 初始化词典对象
void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。


示例：

输入：
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
输出：
[null,null,null,null,false,true,true,true]

解释：
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True


提示：

1 <= word.length <= 500
addWord 中的 word 由小写英文字母组成
search 中的 word 由 '.' 或小写英文字母组成
最调用多 50000 次 addWord 和 search
*/

package main

import (
	"fmt"
)

type WordDictionary struct {
	end  int
	next map[string]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		end:  0,
		next: map[string]*WordDictionary{},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if word == "" {
		return
	}
	node := this
	length := len(word)
	for i := 0; i < length; i++ {
		s := word[i : i+1]
		if _, ok := node.next[s]; !ok {
			node.next[s] = &WordDictionary{
				end:  0,
				next: map[string]*WordDictionary{},
			}
		}
		node = node.next[s]
	}
	node.end++
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.searchHelper(word, 0)
}

func (this *WordDictionary) searchHelper(word string, l int) bool {
	node := this
	length := len(word)
	//递归，有一个true即为true

	if length-1 == l {
		//递归出口
		if word[l:l+1] == "." {
			for _, v := range node.next {
				if v.end > 0 {
					return true
				}
			}
			return false
		}

		if _, ok := node.next[word[l:l+1]]; ok {
			if node.next[word[l:l+1]].end > 0 {
				return true
			}
		}

		return false
	}

	//递归体

	s := word[l : l+1]
	if s == "." && node.next != nil {
		for _, v := range node.next {
			if v.searchHelper(word, l+1) {
				return true
			}
		}

	} else {
		if _, ok := node.next[s]; !ok {
			return false
		}
		node = node.next[s]
		if node.searchHelper(word, l+1) {
			return true
		}
	}

	return false

}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	root := &WordDictionary{
		end:  0,
		next: map[string]*WordDictionary{},
	}
	//s := "abc."
	//fmt.Println(s[3])
	//fmt.Println(reflect.TypeOf(s[3]))
	root.AddWord("bad")
	root.AddWord("dad")
	root.AddWord("add")
	//root.AddWord("addWord")
	fmt.Println(root.Search("pad"))
	fmt.Println(root.Search("bad"))
	fmt.Println(root.Search(".ad"))
	fmt.Println(root.Search("..d"))
	fmt.Println(root.Search("a.d."))
}
