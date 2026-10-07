package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	k := 1

	sum := maximumSubarraySum(nums, k)
	fmt.Println(sum)
}

func maximumSubarraySum(nums []int, k int) int64 {
	prefixSum := int64(0)
	result := int64(0)
	found := false
	minPrefixByValue := make(map[int64]int64)

	for _, num := range nums {
		value := int64(num)
		endSum := prefixSum + value

		for _, startValue := range []int64{value - int64(k), value + int64(k)} {
			if minPrefix, ok := minPrefixByValue[startValue]; ok {
				sum := endSum - minPrefix
				if !found || sum > result {
					result = sum
					found = true
				}
			}
		}

		if minPrefix, ok := minPrefixByValue[value]; !ok || prefixSum < minPrefix {
			minPrefixByValue[value] = prefixSum
		}

		prefixSum = endSum
	}

	if !found {
		return 0
	}

	return result
}

func reduce(acc []int64, f func(int64, int64) int64) int64 {
	if len(acc) == 0 {
		return 0
	}

	result := acc[0]

	for i := 1; i < len(acc); i++ {
		result = f(result, acc[i])
	}

	return result
}
