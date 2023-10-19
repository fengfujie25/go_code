package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

// 反转链表函数
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

// func main() {
//     // 创建链表
//     head := &ListNode{Val: 1, Next: nil}
//     node1 := &ListNode{Val: 2, Next: nil}
//     node2 := &ListNode{Val: 3, Next: nil}
//     head.Next = node1
//     node1.Next = node2

//     // 反转链表并打印结果
//     newHead := reverseList(head)
//     for newHead != nil {
//         fmt.Println(newHead.Val)
//         newHead = newHead.Next
//     }
// }


func maxProfit(prices []int) int {
    maxProfit := 0        // 设置初始的最大利润为0

    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            maxProfit += prices[i] - prices[i-1] // 计算每天的利润并累加
        }
    }

    return maxProfit
}

func main() {
    prices := []int{7, 1, 5, 3, 6, 4}
    fmt.Println(maxProfit(prices)) // 输出：7
}