package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	l := &ListNode{}
	current := l
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return l.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		mergedList := make([]*ListNode, 0)
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				mergedList = append(mergedList, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				mergedList = append(mergedList, lists[i])
			}
		}
		lists = mergedList
	}

	return lists[0]
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func main() {
	printList(mergeKLists([]*ListNode{
		createList([]int{1, 4, 5}),
		createList([]int{1, 3, 4}),
		createList([]int{2, 6}),
	}))

	printList(mergeKLists([]*ListNode{}))

	printList(mergeKLists([]*ListNode{
		createList([]int{}),
	}))
}
