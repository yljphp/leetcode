package main

import "fmt"

/**
面试题57 - II. 和为s的连续正数序列
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。



示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]


限制：

1 <= target <= 10^5
@see https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
 */
func main() {

	fmt.Println(findContinuousSequence(15))
}

func findContinuousSequence(target int) [][]int {
	var coi [][]int
	l, r := 1, 2
	for l < r {
		//等差数列前n项和公式
		sum := (l + r) * (r - l + 1) / 2
		if sum == target {
			var lin []int
			for i := l; i <= r; i++ {
				lin = append(lin, i)
			}
			coi = append(coi, lin)
			l++
		} else {
			if sum < target {
				r++
			} else {
				l++
			}
		}
	}
	return coi
}

func findContinuousSequenceBak(target int) [][]int {

	j := 0

	resInfo := make([][]int, 0)

	for {

		tmpRes := make([]int, 0)

		tmpNum := 0

		for i := j + 1; i < target; i++ {

			tmpNum += i

			tmpRes = append(tmpRes, i)

			if tmpNum > target {
				break
			} else if tmpNum == target {
				resInfo = append(resInfo, tmpRes)
				break
			} else {
				continue
			}
		}

		j++

		if j >= target {
			break
		}
	}

	return resInfo

}
