// Time Complexity: o(n)
// Space Complexity: o(1)
// Runtime: 0 ms
// Memory: 4.2 MB

// Key Idea:
// 1. two pointers (runner, stopper)
// 2. move runner n steps ahead
// 3. move both pointers until runner reaches the end
// 4. remove the nth node from the end by adjusting the stopper's Next pointer

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    runner := head
    stopper := head

    // Move the runner n steps ahead
    for i := 0; i < n && runner != nil; i++ {
        runner = runner.Next
    }

    if runner == nil {
        // If n is equal to the length of the list, remove the head
        return head.Next 
    }

    for head != nil  {
        if runner.Next == nil {
            stopper.Next = stopper.Next.Next
            return head   
        }
            
        runner = runner.Next
        stopper = stopper.Next
    }

    return head
}