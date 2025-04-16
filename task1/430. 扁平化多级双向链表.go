/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
    current := root // 初始化current指针为链表头节点
    for current != nil { // 遍历链表直到末尾
        if current.Child != nil { // 如果当前节点有子链表
            nextTemp := current.Next // 保存当前节点的原始下一个节点
            childHead := current.Child // 获取子链表的头节点
            childTail := childHead// 寻找子链表的尾节点
            for childTail.Next != nil {
                childTail = childTail.Next
            }
            // 将子链表插入到当前节点之后
            current.Next = childHead
            childHead.Prev = current
            // 将子链表的尾节点连接到原链表的下一个节点
            childTail.Next = nextTemp
            if nextTemp != nil {
                nextTemp.Prev = childTail
            }
            current.Child = nil // 清空当前节点的子指针
        }
        current = current.Next // 移动到下一个节点继续处理
    }
    return root // 返回扁平化后的头节点
}
