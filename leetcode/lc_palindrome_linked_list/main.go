/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    for head.Next != nil {
        tail := head.Next
        temp := head
        for tail.Next!=nil {
            temp = tail
            tail = tail.Next
        }
        if head.Val != tail.Val {
            return false
        }
        head = head.Next
        temp.Next = nil
    }
    return true
}

func isPalindrome(head *ListNode) bool {
    str := ""
    for head!=nil {
        str += string(head.Val)
        head = head.Next
    }
    for i:=0; i<len(str)/2; i++ {
        if str[i] != str[len(str)-1-i] {
            return false
        }
    }
    return true
}

---------------------------------------------------
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    fast := head
    slow := head
    temp := slow

    for fast!=nil && fast.Next!=nil {
        temp = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    slow = reverse(slow)
    temp.Next = slow
    temp = slow

    for head != temp {
        if head.Val != slow.Val {
            return false
        }
        slow = slow.Next
        head = head.Next
    }

    return true
}

func reverse(n *ListNode) *ListNode {
    curr := n.Next
    prev := n
    prev.Next = nil

    for curr!=nil {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }

    return prev
}

---------------------------------------------------
func isPalindrome(head *ListNode) bool {
    arr := make([]int, 0)

    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }

    for i:=0; i<len(arr); i++ {
        if arr[i] != arr[len(arr)-1-i] {
            return false
        }
    }
    return true
}

---------------------------------------------------

func isPalindrome(head *ListNode) bool {
    // reverse till the end pointer points to the end of the list
    // start comparing the nodes
    //      return false if the nodes don't match
    
    end := head
    var prev *ListNode = nil
    
    // even - end is nil - curr vs prev
    // odd - end.Next is nil - skip 1 node then curr vs prev
    // 0<-1<-2->3->4->5
    //  -- --p--h-- --e
    for end != nil && end.Next != nil {
        end = end.Next.Next
        first := head.Next
        head.Next = prev
        prev = head
        head = first
    }
    
    if end == nil {
        // it's even
    } else {
        // it's odd
        // skip one node
        head = head.Next
    }
    
    // compare head and previous
    for head != nil && prev != nil {
        if head.Val != prev.Val {
            return false
        }
        head = head.Next
        prev = prev.Next
    }
    
    return true
}
