/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/9
 * Time 1:58 下午
 */

package main

import (
	"fmt"
	"sort"
)

/**
给出 字符串 text 和 字符串列表 words, 返回所有的索引对 [i, j] 使得在索引对范围内的子字符串 text[i]...text[j]（包括 i 和 j）属于字符串列表 words。



示例 1:

输入: text = "thestoryofleetcodeandme", words = ["story","fleet","leetcode"]
输出: [[3,7],[9,13],[10,17]]
示例 2:

输入: text = "ababa", words = ["aba","ab"]
输出: [[0,1],[0,2],[2,3],[2,4]]
解释:
注意，返回的配对可以有交叉，比如，"aba" 既在 [0,2] 中也在 [2,4] 中


提示:

所有字符串都只包含小写字母。
保证 words 中的字符串无重复。
1 <= text.length <= 100
1 <= words.length <= 20
1 <= words[i].length <= 50
按序返回索引对 [i,j]（即，按照索引对的第一个索引进行排序，当第一个索引对相同时按照第二个索引对排序）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/index-pairs-of-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func indexPairs(text string, words []string) [][]int {
	// 数据量小，直接暴力求解
	lT := len(text)
	lA := len(words)
	s := make([][]int, 0)
	for i := 0; i < lA; i++ {
		lW := len(words[i])
		for j := 0; j+lW <= lT; j++ {
			if words[i] == text[j:j+lW] {
				s = append(s, []int{j, j + lW - 1})
			}
		}
	}
	sort.Slice(s, func(i, j int) bool {
		if s[i][0] == s[j][0] {
			return s[i][1] < s[j][1]
		} else {
			return s[i][0] < s[j][0]
		}

	})
	return s
}

func main() {
	text := "ababa"
	words := []string{"aba", "ab"}
	fmt.Println(indexPairs(text, words))
}
