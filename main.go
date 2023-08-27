package main

import (
	"log"
	"sort"
)

func main() {
	log.Println(MaxProfit([]int{31, 94, 91, 92, 4, 11, 2, 20, 59, 80, 1, 3, 2, 7}, 3))
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println(ShiftArray(array, -1))
}

func MaxProfit(prices []int, i int) int {
	if i == 0 {
		log.Println("Jumlah transaksi harus minimal 1")
		return -1
	}
	if len(prices) < 2 {
		return 0
	}
	var lowest int = -1
	var highest int = -1
	var profits []int
	for x := 0; x < len(prices); x++ {
		// if previous price is higher than current prices, set profits
		if x != 0 && prices[x-1] > prices[x] {
			if highest-lowest > 0 {
				// if the previous price is higher, set the profits value to array
				profits = append(profits, highest-lowest)
			}
			highest = -1
			lowest = prices[x]
		} else if x == 0 {
			// if still the first price, set lowest
			lowest = prices[x]

		} else if lowest != -1 && prices[x-1] < prices[x] {
			// if current price is the end of the array and the current price is higher
			highest = prices[x]
		} else if lowest != -1 && x == len(prices)-1 && prices[x-1] < prices[x] {
			// if current price is the end of the array and the current price is higher
			highest = prices[x]
		}
		if x == len(prices)-1 {
			// at the end of the loop, set the profit
			if highest-lowest > 0 {
				profits = append(profits, highest-lowest)
			}
		}
	}
	sort.Slice(profits, func(x, y int) bool {
		return profits[x] > profits[y]
	})
	var nums int = 0
	for z := 0; z < i; z++ {
		if z < len(profits) {
			nums += profits[z]
		}
	}
	return nums
}

func ShiftArray(array []int, i int) []int {
	if i < 0 {
		log.Println("Jumlah perputaran array harus bernilai positif")
		return []int{}
	}
	t := [9]int{}
	for x := 0; x < i; x++ {
		t[0] = array[3]
		t[1] = array[0]
		t[2] = array[1]
		t[3] = array[6]
		t[4] = array[4]
		t[5] = array[2]
		t[6] = array[7]
		t[7] = array[8]
		t[8] = array[5]
		copy(array, t[:])
	}

	return array
}
