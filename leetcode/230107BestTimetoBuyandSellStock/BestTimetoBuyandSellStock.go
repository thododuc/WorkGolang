package main //https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan&id=level-1
import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	buy, sell := prices[0], prices[0]
	max := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			sell = buy
			if sell-buy > max {
				max = sell - buy
			}
		} else {
			if prices[i] > sell {
				sell = prices[i]
				if sell-buy > max {
					max = sell - buy
				}
			}
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
