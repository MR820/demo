/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/10
 * Time 1:36 下午
 */

/**
在英语中，我们有一个叫做 词根(root)的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。

现在，给定一个由许多词根组成的词典和一个句子。你需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。

你需要输出替换之后的句子。



示例 1：

输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"
示例 2：

输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
输出："a a b c"
示例 3：

输入：dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
输出："a a a a a a a a bbb baba a"
示例 4：

输入：dictionary = ["catt","cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"
示例 5：

输入：dictionary = ["ac","ab"], sentence = "it is abnormal that this solution is accepted"
输出："it is ab that this solution is ac"


提示：

1 <= dictionary.length <= 1000
1 <= dictionary[i].length <= 100
dictionary[i] 仅由小写字母组成。
1 <= sentence.length <= 10^6
sentence 仅由小写字母和空格组成。
sentence 中单词的总量在范围 [1, 1000] 内。
sentence 中每个单词的长度在范围 [1, 1000] 内。
sentence 中单词之间由一个空格隔开。
sentence 没有前导或尾随空格。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/replace-words
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"strings"
)

type Node struct {
	end  int
	next map[string]*Node
}

func (root *Node) insert(word string) {
	if word == "" {
		return
	}
	node := root
	length := len(word)
	for i := 0; i < length; i++ {
		c := word[i : i+1]
		if _, ok := node.next[c]; !ok {
			//不包含
			node.next[c] = &Node{
				end:  0,
				next: map[string]*Node{},
			}
		}
		node = node.next[c]
	}
	node.end++
}

func (root *Node) prefix(s string) string {
	//s的最短前缀
	node := root
	l := len(s)
	for i := 0; i < l; i++ {
		if _, ok := node.next[s[i:i+1]]; !ok {
			return s
		}
		node = node.next[s[i:i+1]]

		if node.end > 0 {
			return s[0 : i+1]
		}
	}
	return s
}

func replaceWords(dictionary []string, sentence string) string {
	// 把词根放入前缀树，在树上查找每个单词的最短词根
	length := len(dictionary)
	root := &Node{next: map[string]*Node{}}
	for i := 0; i < length; i++ {
		root.insert(dictionary[i])
	}

	s := strings.Split(sentence, " ")

	l := len(s)
	for i := 0; i < l; i++ {
		res := root.prefix(s[i])
		s[i] = res
	}
	return strings.Join(s, " ")
}

func main() {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence))
}
