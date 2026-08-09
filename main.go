package main

import (
	"fmt"
)

// TODO: bubbleSort без использования built-in sort
func bubbleSort(nums []int) []int {
	// Создаём копию среза, чтобы не менять оригинал
	result := make([]int, len(nums))
	copy(result, nums)

	n := len(result)
	if n <= 1 {
		return result
	}

	swapped := true
	for i := 0; i < n-1 && swapped; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				// Меняем местами
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
	}

	return result
}



func main() {
    fmt.Println(bubbleSort([]int{5, 3, 1, 4, 2})) // → [1, 2, 3, 4, 5]
    fmt.Println(bubbleSort([]int{10, 9, 8}))       // → [8, 9, 10]
    fmt.Println(bubbleSort([]int{1}))               // → [1]      // → 0
}