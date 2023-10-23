// This solution does not give the index og the original array, as the array is now sorted inplace.
func twoSum(nums []int, target int) []int {
    sort.Ints(nums)
    for i,j := 0,  len(nums) - 1; i < len(nums) && j >= i;  {
        if nums[j] + nums[i] == target {
            return []int{nums[j],nums[i]}   
        } else if nums[j] + nums[i] > target {
            j--
        } else {
            i++
        }
    }
    return []int{-1,-1} 
}

//Time Complexity : O(n)
No extra space complexity
