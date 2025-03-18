// Time Complexity: o(n)
// Space Complexity: o(n)
// Runtime: 0 ms
// Memory: 4.66 MB

// Key Idea:
// 1. stack
// 2. reverced head (at last node)
// 3. change every node Next value by stack

// optimization
// 1. in-place reversal by using 3 pointers (prev, curr, next)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	stack := []*ListNode{}

	// push node pointer to stack
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	// error handle , if stack is empty
	if len(stack) == 0 {
		return nil
	}

	// get reverced head
	revercedHead := stack[len(stack)-1]

	// set current head at revercedHead
	curr := revercedHead

	// pop last value
	stack = stack[:len(stack)-1]

	// linked list traversal
	for len(stack) > 0 {
		// change addr to next node
		curr.Next = stack[len(stack)-1]
		// pop stack
		stack = stack[:len(stack)-1]
		// renew current node addr
		curr = curr.Next
	}

	curr.Next = nil
	return revercedHead
}