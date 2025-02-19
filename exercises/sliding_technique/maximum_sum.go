package slidingtechnique

// Problem: Given an array of integers nums and a positive integer k, find the maximum sum of any contiguous subarray (window) of size k.

func MaxSum(numbers []int, k int) int {
  left := 0
  right := 0
  maxSum := 0 
  sum := 0

  for right < len(numbers){
    sum += numbers[right]
    if (right - left) == k-1{
      if sum > maxSum{
        maxSum = sum
      }
      sum -= numbers[left]
      left++
    }
    right++
  } 

  return maxSum
}
