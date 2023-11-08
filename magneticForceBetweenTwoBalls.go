func maxDistance(position []int, m int) int {
  sort.Ints(position)
    
    countBalls := func(minDist int) int {
      count := 1
      prev := position[0]
      for _, p := range position {
        if (p-prev) >= minDist {
          count ++
          prev = p
        }
      }
      return count
    }
    left := position[0]
    right := position[len(position)-1] - position[0]
    for (left < right) {
      mid := left + (right-left+1)/2
      if countBalls(mid) >= m {      
        left = mid
      } else {
        right = mid - 1
      }
    }
    return left
}
