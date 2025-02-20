package slidingtechnique

// Given an array of integers nums and a positive integer k, determine if there are any two distinct indices i and j in the array where nums[i] == nums[j] and abs(i - j) <= k. (In other words, are there any duplicates within a window of size k?)

func ContainsDuplicate(arr []int, k int) bool {
  left := 0
  right := 0

  for right < len(arr) -1 {
    right++
    if right-left > k -1 {
      left++
    }
    if left != right && arr[left] == arr[right] {
      return true
    }   
  }
  return false
}
