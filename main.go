package main

import (
	"log"
	"sort"
)

func main() {
	log.Println(MaxProfit([]int{4, 11, 2, 20, 59, 80, 1, 20, 90}, 2))
}

func MaxProfit(prices []int, i int) int {
	if i == 0 {
		log.Fatalln("Jumlah transaksi harus minimal 1")
	}
	if len(prices) < 2 {
		return 0
	}
	var lowest int = 0
	var highest int = 0
	var profits []int
	for x := 0; x < len(prices)-1; x++ {
		if prices[x] <= prices[x+1] {
			if x != 0 && prices[x-1] > prices[x] {
				profits = append(profits, highest-lowest)
			}
			if x == 0 || prices[x-1] > prices[x] {
				lowest = prices[x]
			}
			if x+1 == len(prices)-1 && prices[x] <= prices[x+1] {
				highest = prices[x+1]
				profits = append(profits, highest-lowest)
				break
			}
			continue
		} else {
			highest = prices[x]
		}
	}
	sort.Slice(profits, func(x, y int) bool {
		return profits[x] > profits[y]
	})
	var nums int = 0
	for z := 0; z < i; z++ {
		nums += profits[z]
	}
	return nums
}
