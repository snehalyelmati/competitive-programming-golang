/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    hmap := make(map[*ListNode]bool)
    for head!=nil {
        if _, ok := hmap[head]; ok {
            return true
        } else {
            hmap[head] = true
        }
        head = head.Next
    }
    return false
}

-----------------------------------
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    fast := head.Next
    slow := head

    for fast != slow {
        if fast.Next == nil || fast.Next.Next == nil {
            return false
        }
        if fast == slow {
            break
        }
        fast = fast.Next.Next
        slow = slow.Next
    }

    return true
}
