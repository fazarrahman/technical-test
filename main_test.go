package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxProfit_IncorrectTransactionValueError_ReturnMinusOne(t *testing.T) {
	res := MaxProfit([]int{31, 94, 91, 92, 4, 11, 2, 20, 59, 80, 1, 3, 2, 7}, 0)
	assert.Equal(t, -1, res)
}

func Test_MaxProfit_ArrayInput1(t *testing.T) {
	res := MaxProfit([]int{4, 11, 2, 20, 59, 80}, 2)
	assert.Equal(t, 85, res)
}

func Test_MaxProfit_PricesLengthIsOnlyOne_ReturnZero(t *testing.T) {
	res := MaxProfit([]int{11}, 2)
	assert.Equal(t, 0, res)
}

func Test_MaxProfit_ArrayInput2(t *testing.T) {
	res := MaxProfit([]int{31, 94, 91, 92, 4, 11, 2, 20, 59, 80, 1, 3, 2, 7}, 3)
	assert.Equal(t, 148, res)
}

func Test_ShiftArray_ShiftMinusOneTime(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := ShiftArray(array, -1)
	assert.Equal(t, []int{}, res)
}

func Test_ShiftArray_ShiftZeroTime(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := ShiftArray(array, 0)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)
}

func Test_ShiftArray_ShiftOnce(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := ShiftArray(array, 1)
	assert.Equal(t, []int{4, 1, 2, 7, 5, 3, 8, 9, 6}, res)
}

func Test_ShiftArray_ShiftTwice(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := ShiftArray(array, 2)
	assert.Equal(t, []int{7, 4, 1, 8, 5, 2, 9, 6, 3}, res)
}
