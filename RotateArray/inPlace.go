func rotate(nums []int, k int)  {       
    // First rotate the whole array
    start, end := 0, len(nums)-1
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
    //Rotate the first k elements
    if k %= len(nums); k > 0 {
        start, end = 0, k-1
        for start < end {
            nums[start], nums[end] = nums[end], nums[start]
            start++
            end--
        }
    }
    //Rotate the last k elements 
    start,end = k, len(nums)-1
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
