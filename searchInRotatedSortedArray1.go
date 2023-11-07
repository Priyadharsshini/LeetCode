// https://leetcode.com/problems/search-in-rotated-sorted-array/description/

func search(nums []int, target int) int {
    start := 0
    end := len(nums)-1
    for (start < end){
        mid := start +(end-start)/2
        if nums[mid] == target {
            return mid
        }
        if (nums[start] <= nums[mid]){
            if (target >= nums[start] && target <= nums[mid]){
                end = mid 
            } else {
                start = mid 
            }
        } else {
            if (target > nums[mid] && target <= nums[end]){
                start = mid + 1
            } else {
                end = mid - 1
            }
        }

    }
    return -1
}
