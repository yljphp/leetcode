package main

import "fmt"

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
