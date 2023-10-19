package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 写一个链表反转函数
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func TestReverseList(t *testing.T) {
	// Test case 1: empty list
	list1 := (*ListNode)(nil)
	reversed1 := reverseList(list1)
	assert.Nil(t, reversed1)

	// Test case 2: single node list
	list2 := &ListNode{Val: 1, Next: nil}
	reversed2 := reverseList(list2)
	assert.Equal(t, 1, reversed2.Val)
	assert.Nil(t, reversed2.Next)

	// Test case 3: multiple node list
	list3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	reversed3 := reverseList(list3)
	assert.Equal(t, 3, reversed3.Val)
	assert.Equal(t, 2, reversed3.Next.Val)
	assert.Equal(t, 1, reversed3.Next.Next.Val)
	assert.Nil(t, reversed3.Next.Next.Next)
}

// 运行测试代码
func main() {
	head := TestReverseList([]int{1, 2, 3, 4, 5})
	fmt.Println(reverseList(head))
}