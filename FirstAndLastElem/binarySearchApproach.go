package main

import "fmt"

func binarySearchFirst(a []int, target int) int {
	left := 0
	right := len(a) - 1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] == target {
			result = mid
			right = mid - 1
		} else if a[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func binarySearchLast(a []int, target int) int {
	left := 0
	right := len(a) - 1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] == target {
			result = mid
			left = mid + 1
		} else if a[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 2, 2, 3, 3, 4, 4, 4, 4, 5}
	firstresult := binarySearchFirst(a, 3)
	lastresult := binarySearchLast(a, 3)
	fmt.Println([]int{firstresult, lastresult})
}
