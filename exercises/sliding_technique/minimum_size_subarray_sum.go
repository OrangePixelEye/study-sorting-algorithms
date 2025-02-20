package slidingtechnique


// Given an array of positive integers nums and a positive integer target, find the minimal length of a contiguous subarray whose sum is greater than or equal to target. If there is no such subarray, return 0

func MinumumSizeSubarray(nums[]int, target int) int {
	minLength := int(^uint(0) >> 1) 
	windowStart := 0
	windowSum := 0

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		windowSum += nums[windowEnd]

		for windowSum >= target {
			currentLength := windowEnd - windowStart + 1
			if currentLength < minLength {
				minLength = currentLength
			}
			windowSum -= nums[windowStart]
			windowStart++
		}
	}

	if minLength == int(^uint(0)>>1) {
		return 0
	}

	return minLength
}
