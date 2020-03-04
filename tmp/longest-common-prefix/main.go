package main

import (
	"strings"
)

/**

14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
@see https://leetcode-cn.com/problems/longest-common-prefix/
*/
func main() {

	tmp := []string{
		"flower", "flow", "flight",
	}

	longestCommonPrefix(tmp)

}


func longestCommonPrefix(strs []string) string {

	if len(strs) <= 0 {
		return ""
	}

	//以第一个字符串为最大公共前缀
	firstStr := strs[0]

	//从第二个字符串开始判断是否存在该前缀
	for i := 1; i < len(strs); i++ {

		for  {

			//fmt.Println(strs[i], firstStr)
			//flow flower
			//flow flowe
			//flow flow
			//flight flow
			//flight flo
			//flight flx
			//fl
			///不存在时将字符串从后开始缩减直到存在，然后挨个遍历字符串数组。

			//一定要是0 不能是-1 比如遇到["c","acc","bcc"]情况
			if strings.Index(strs[i], firstStr) !=0 {

				firstStr = firstStr[0 : len(firstStr)-1]
				continue
			}

			break
		}

	}

	return firstStr
}