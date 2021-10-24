/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/11
 * Time 10:51 上午
 */

/**
1023. 驼峰式匹配
如果我们可以将小写字母插入模式串 pattern 得到待查询项 query，那么待查询项与给定模式串匹配。（我们可以在任何位置插入每个字符，也可以插入 0 个字符。）

给定待查询列表 queries，和模式串 pattern，返回由布尔值组成的答案列表 answer。只有在待查项 queries[i] 与模式串 pattern 匹配时， answer[i] 才为 true，否则为 false。



示例 1：

输入：queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FB"
输出：[true,false,true,true,false]
示例：
"FooBar" 可以这样生成："F" + "oo" + "B" + "ar"。
"FootBall" 可以这样生成："F" + "oot" + "B" + "all".
"FrameBuffer" 可以这样生成："F" + "rame" + "B" + "uffer".
示例 2：

输入：queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBa"
输出：[true,false,true,false,false]
解释：
"FooBar" 可以这样生成："Fo" + "o" + "Ba" + "r".
"FootBall" 可以这样生成："Fo" + "ot" + "Ba" + "ll".
示例 3：

输出：queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBaT"
输入：[false,true,false,false,false]
解释：
"FooBarTest" 可以这样生成："Fo" + "o" + "Ba" + "r" + "T" + "est".


提示：

1 <= queries.length <= 100
1 <= queries[i].length <= 100
1 <= pattern.length <= 100
所有字符串都仅由大写和小写英文字母组成。

*/
package main

import (
	"fmt"
)

type Node struct {
	end  int
	next map[int32]*Node
}

func Constructor() *Node {
	return &Node{
		end:  0,
		next: map[int32]*Node{},
	}
}

func (this *Node) Insert(word string) {
	node := this
	for _, v := range word {
		if _, ok := node.next[v]; !ok {
			node.next[v] = Constructor()
		}
		node = node.next[v]
	}
	node.end++
}

func (this *Node) Match(word string) bool {
	node := this
	for _, v := range word {
		if _, ok := node.next[v]; !ok {
			if v < 97 {
				//大写字母
				return false
			}
			//小写字母直接pass
		} else {
			node = node.next[v]
		}
	}
	if node.end > 0 {
		return true
	}
	return false
}

func camelMatch(queries []string, pattern string) []bool {
	//pattern生成字典树
	//字符串去匹配，如果大写字母不匹配，直接返回false
	//小写字母不匹配，忽略。小写字母匹配，字典树节点下移
	//一直将字符串遍历到最后 O(M+N)
	// a 97 A 65
	root := Constructor()
	root.Insert(pattern)
	b := make([]bool, 0)
	for _, v := range queries {
		res := root.Match(v)
		b = append(b, res)
	}
	return b

}

func main() {
	querys := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FB"
	fmt.Print(camelMatch(querys, pattern))
}
