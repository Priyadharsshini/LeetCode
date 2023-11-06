//https://leetcode.com/problems/peak-index-in-a-mountain-array/

func peakIndexInMountainArray(arr []int) int {
    start := 0
    end := len(arr) - 1
    mid := 0
    for start < end {
        mid = start + (end-start)/2
        if (arr[mid] < arr[mid+1]){
            start = mid+1
        } else {
            end = mid
        }

    }
     return start
}
