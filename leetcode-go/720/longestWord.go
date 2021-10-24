/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/9
 * Time 5:54 下午
 */

/**
720. 词典中最长的单词
给出一个字符串数组words组成的一本英语词典。从中找出最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。若其中有多个可行的答案，则返回答案中字典序最小的单词。

若无答案，则返回空字符串。



示例 1：

输入：
words = ["w","wo","wor","worl", "world"]
输出："world"
解释：
单词"world"可由"w", "wo", "wor", 和 "worl"添加一个字母组成。
示例 2：

输入：
words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
输出："apple"
解释：
"apply"和"apple"都能由词典中的单词组成。但是"apple"的字典序小于"apply"。


提示：

所有输入的字符串都只包含小写字母。
words数组长度范围为[1,1000]。
words[i]的长度范围为[1,30]。
*/

package main

import "fmt"

type Node struct {
	pass int
	end  int
	word string
	next [26]*Node
}

func Constructor() *Node {
	return &Node{
		pass: 0,
		end:  0,
		word: "",
		next: [26]*Node{},
	}
}

func (root *Node) insert(word string) {
	node := root
	for _, v := range word {
		index := v - 'a'
		if node.next[index] == nil {
			node.next[index] = Constructor()
		}
		node.next[index].word = node.word + string(v)
		node = node.next[index]
		node.pass++
	}
	node.end = 1
}

func longestWord(words []string) string {
	//前缀树,遍历所有节点。end!=0的从树中脱离
	//深度遍历，找出最大长度（相同长度以第一个为最大）

	//var s string
	root := Constructor()
	//构造前缀树
	for _, v := range words {
		root.insert(v)
	}

	return root.dfs()
}

/**
深度优先遍历，适合查找最优解
*/
func (this *Node) dfs() string {
	s := make([]*Node, 0)
	str := ""
	s = append(s, this)
	for len(s) > 0 {
		l := len(s)
		node := s[l-1]
		s = s[:l-1]
		for _, v := range node.next {
			if v == nil {
				continue
			}
			if v.end > 0 {
				if v.pass+len(v.word) <= len(str) {
					//剪枝
					break
				}
				if len(v.word) > len(str) {
					str = v.word
				} else if len(v.word) == len(str) && v.word < str {
					str = v.word
				}
				s = append(s, v)
			}
		}

	}
	return str
}
func main() {
	words := []string{"ovvnb",
		"itcudluaycsspjrohf",
		"cqwec",
		"nz",
		"mkomsvmrauphvm",
		"ix"}
	fmt.Println(longestWord(words))

}
