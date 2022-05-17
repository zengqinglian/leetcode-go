package main

/*
Runtime: 14 ms, faster than 63.43% of Go online submissions for Two Sum II - Input Array Is Sorted.
Memory Usage: 5.3 MB, less than 91.90% of Go online submissions for Two Sum II - Input Array Is Sorted.
*/
func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return nil
}
