package main

import (
	"fmt"
	"time"
)

func min(item []int) (int, int) {
	i := 0
	result := item[0]
	position := 0
	for i < len(item) {
		candidate := item[i]
		if result > candidate {
			result = candidate
			position = i
		}
		i++
	}

	return result, position
}

func swap(item []int, p1 int, p2 int) []int {
	r2 := item[p1]
	item[p1] = item[p2]
	item[p2] = r2
	return item
}

func SelectionSort(item []int) []int {
	i := 0
	for i < len(item)-1 {
		_, min := min(item[i:])
		min = min + i
		if min != 0 {
			swap(item, i, min)
		}
		i++
	}

	return item
}

func InsertionSort(item []int) []int {
	for i := 0; i < len(item); i++ {
		if i != 0 {
			for u := 0; u < i && item[i] < item[u]; u++ {
				swap(item, i, u)
			}
		}
	}
	return item
}

func GnomeSort(item []int) []int {
	i := 0
	for i < len(item) {
		if i == 0 || item[i] >= item[i-1] {
			i++
		} else {
			item[i], item[i-1] = item[i-1], item[i]
			i--
		}
	}
	return item
}

func BubbleSort(item []int) []int {
	for i := 0; i < len(item)-1; i++ {
		for u := i + 1; u < len(item); u++ {
			if item[i] > item[u] {
				swap(item, i, u)
			}
		}
	}
	return item
}

func CocktailShakerSort(item []int) []int {
	n := len(item)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if item[i] > item[i+1] {
				item = swap(item, i, i+1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
		swapped = false
		for i := n - 2; i >= 0; i-- {
			if item[i] > item[i+1] {
				item = swap(item, i, i+1)
				swapped = true
			}
		}
	}
	return item
}

func getMax(item []int) int {
	max := item[0]
	for _, num := range item {
		if num > max {
			max = num
		}
	}
	return max
}

func countSort(item []int, exp int) {
	n := len(item)
	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		count[(item[i]/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		output[count[(item[i]/exp)%10]-1] = item[i]
		count[(item[i]/exp)%10]--
	}

	for i := 0; i < n; i++ {
		item[i] = output[i]
	}
}

func RadixSort(item []int) {
	max := getMax(item)

	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(item, exp)
	}
}

func main() {

	item := []int{170, 45, 75, 90, 802, 24, 2, 66}
	start := time.Now()
	RadixSort(item)
	runtime := time.Since(start)
	fmt.Println(item)
	fmt.Println(runtime)

}
