package main

import (
	"fmt"
	"sort"
)

/**
面试题 10.01. 合并排序的数组
给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。

初始化 A 和 B 的元素数量分别为 m 和 n。

示例:

输入:
A = [1,2,3,0,0,0], m = 3
B = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
@see https://leetcode-cn.com/problems/sorted-merge-lcci/
 */
func merge(A []int, m int, B []int, n int)  {


	tmp := append(A[0:m], B[0:n]...)


	sort.Ints(tmp)

	fmt.Println(tmp)

}
//
//func reverse(s []int) []int {
//	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
//		s[i], s[j] = s[j], s[i]
//	}
//	return s
//}

func main() {

	merge([]int{1,2,3,0,0,0},3,[]int{2,5,6},3)
}
