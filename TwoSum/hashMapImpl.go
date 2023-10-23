func twoSum(nums []int, target int) []int {
    hm := make(map[int]int)
    for index, value := range nums {
      // Subtract the current array element from the target element and check if that resulting value is in the hasmap.
      // If present then return both index
        currentValue := target - value
        if val, ok := hm[currentValue];ok {
            return []int{val, index}
        } else {
        // If the resulting value is not in hashmap, then store that value along with the index in hashmap
            hm[value] = index
        }
    }
    return []int{-1,-1}    
}

//Time Complexity : O(n) as we are iterating through the whole original array once in the worst case
//Space Complexity : O(n) as we are using the hashmap as an additional storage.
