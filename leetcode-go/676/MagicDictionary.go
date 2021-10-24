/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/13
 * Time 1:07 下午
 */

package main

import (
	"fmt"
)

/**
676. 实现一个魔法字典
设计一个使用单词列表进行初始化的数据结构，单词列表中的单词 互不相同 。 如果给出一个单词，请判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单词存在于你构建的字典中。

实现 MagicDictionary 类：

MagicDictionary() 初始化对象
void buildDict(String[] dictionary) 使用字符串数组 dictionary 设定该数据结构，dictionary 中的字符串互不相同
bool search(String searchWord) 给定一个字符串 searchWord ，判定能否只将字符串中 一个 字母换成另一个字母，使得所形成的新字符串能够与字典中的任一字符串匹配。如果可以，返回 true ；否则，返回 false 。


示例：

输入
["MagicDictionary", "buildDict", "search", "search", "search", "search"]
[[], [["hello", "leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
输出
[null, null, false, true, false, false]

解释
MagicDictionary magicDictionary = new MagicDictionary();
magicDictionary.buildDict(["hello", "leetcode"]);
magicDictionary.search("hello"); // 返回 False
magicDictionary.search("hhllo"); // 将第二个 'h' 替换为 'e' 可以匹配 "hello" ，所以返回 True
magicDictionary.search("hell"); // 返回 False
magicDictionary.search("leetcoded"); // 返回 False


提示：

1 <= dictionary.length <= 100
1 <= dictionary[i].length <= 100
dictionary[i] 仅由小写英文字母组成
dictionary 中的所有字符串 互不相同
1 <= searchWord.length <= 100
searchWord 仅由小写英文字母组成
buildDict 仅在 search 之前调用一次
最多调用 100 次 search

*/

type MagicDictionary struct {
	end  int
	next [26]*MagicDictionary
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{
		end:  0,
		next: [26]*MagicDictionary{},
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {

	for _, v := range dictionary {
		node := this
		for _, val := range v {
			index := val - 'a'
			//fmt.Println(index)
			if node.next[index] == nil {
				root := Constructor()
				node.next[index] = &root
			}
			node = node.next[index]
		}
		node.end++
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return this.searchHelper(searchWord, 0, 0)
}

func (this *MagicDictionary) searchHelper(word string, l int, num int) bool {
	if num > 1 {
		return false
	}
	length := len(word)
	node := this
	//递归出口
	if l == length {
		return num == 0 && node.end == 1
	}

	index := word[l] - 'a'
	//fmt.Println(index)
	for i := 0; i < 26; i++ {
		if node.next[i] == nil {
			continue
		}
		if int(index) == i {
			if node.next[index].searchHelper(word, l+1, num) {
				return true
			}
		} else if node.next[i].searchHelper(word, l+1, num+1) {
			return true
		}
	}
	return false

}

func main() {
	root := Constructor()
	node := &root
	s := []string{"hello", "leetcode"}
	node.BuildDict(s)
	fmt.Println(root.Search("hello"))
	fmt.Println(root.Search("hhllo"))
	fmt.Println(root.Search("hell"))
	fmt.Println(root.Search("leetcoded"))
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
