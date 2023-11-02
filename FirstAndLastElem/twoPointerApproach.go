package main

import "fmt"

func twoPointerApproach(a []int, target int) []int {
	left := 0
	right := len(a) - 1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] == target {
			for left >= 0 && a[left] == target {
				left--
			}
			for right < len(a)-1 && a[right] == target {
				right++
			}
			return []int{left + 1, right - 1}
		} else if a[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{-1, -1}
}

func main() {
	a := []int{1, 2, 2, 2, 3, 3, 4, 4, 4, 4, 5}
	fmt.Println(twoPointerApproach(a, 4))
}
