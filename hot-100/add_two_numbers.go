// Time Complexity: O(Max(M,N))
// Space Complexity: O(Max(M,N))
// Runtime: 0 ms
// Memory: 6.18 MB

// Key Idea:
// 1. Traverse list1 and list2 simultaneously
// 2. Add the values of each node, taking carry into account
// 3. If there's a carry, add it to the value of the next node
// 4. Return the head node of the resulting linked list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create a dummy head for the result list
    dummyHead := &ListNode{}
	// Pointer to the current node in the result list
    current := dummyHead
	// Variable to keep track of carry
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        var v1, v2 int
        if l1 != nil {
			// get the value from list1
            v1 = l1.Val
			// move to the next node in list1
            l1 = l1.Next
        }
        if l2 != nil {
			// get the value from list2
            v2 = l2.Val
			// move to the next node in list2
            l2 = l2.Next
        }

		// Calculate the sum of the two values and the carry
        sum := v1 + v2 + carry
        carry = sum / 10
		// Get the last digit of the sum
        digit := sum % 10

		// Create a new node with the digit and append it to the result list
        current.Next = &ListNode{Val: digit}
		// move to the next node in the result list
        current = current.Next
    }

	// Return the next node of the dummy head, which is the actual head of the result list
    return dummyHead.Next
}