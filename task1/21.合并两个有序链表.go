func mergeTwoLists(list1, list2 *ListNode) *ListNode {
    L1 := ListNode{}
    LL := &L1
    for list1 !=nil && list2 !=nil {
        if list1.Val < list2.Val{
            LL.Next=list1
            list1=list1.Next
        }else{
            LL.Next=list2
            list2=list2.Next
        }
        LL=LL.Next
    }
        if list1 != nil {
        LL.Next = list1
    } else {
        LL.Next = list2
    }
    return L1.Next
}
