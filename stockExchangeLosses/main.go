package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &prices[i])
	}
	maxLoss := 0
	price := prices[0]
	for i := 1; i < len(prices); i++ {
		if price-prices[i] > maxLoss {
			maxLoss = price - prices[i]
		} else if prices[i] > price {
			price = prices[i]
		}
	}
	fmt.Println(0 - maxLoss)
}
