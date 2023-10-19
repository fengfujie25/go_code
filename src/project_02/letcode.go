package project02

//package main

import "fmt"

func maxProfit(prices []int) int {
    minPrice := prices[0] // 设置初始的最低价格为第一天的价格
    maxProfit := 0        // 设置初始的最大利润为0

    for i := 1; i < len(prices); i++ {
        if prices[i] < minPrice {
            minPrice = prices[i] // 更新最低价格
        } else if prices[i]-minPrice > maxProfit {
            maxProfit = prices[i] - minPrice // 更新最大利润
        }
    }

    return maxProfit
}

func main() {
    prices := []int{7, 1, 5, 3, 6, 4}
    fmt.Println(maxProfit(prices)) // 输出：7
}



