package main

import "fmt"

/**
21. 合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

@see https://leetcode-cn.com/problems/merge-two-sorted-lists
 */
func main() {

	l1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				4,
				nil,
			},
		},
	}

	l2 := &ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				4,
				nil,
			},
		},
	}

	mergeTwoLists(l1, l2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {


	var res = &ListNode{}

	node := res

	for {

		if l1.Val <= l2.Val {
			node.Next, l1 = l1, l1.Next
		} else {
			node.Next, l2 = l2, l2.Next
		}

		node = node.Next

		fmt.Println(node.Val)

		if l1 == nil || l2 == nil {
			break
		}
	}

	if l1 != nil {
		node.Next = l1
	} else {
		node.Next = l2
	}

	return res.Next
}
