package main

/*]
Runtime: 28 ms, faster than 92.02% of Go online submissions for Find the Difference of Two Arrays.
Memory Usage: 6.9 MB, less than 92.55% of Go online submissions for Find the Difference of Two Arrays.
*/
func findDifference(nums1 []int, nums2 []int) [][]int {
	var res [][]int
	var empty struct{}
	map_1 := make(map[int]struct{})
	map_2 := make(map[int]struct{})

	for _, v := range nums1 {
		if _, ok := map_1[v]; !ok {
			map_1[v] = empty
		}
	}
	for _, v := range nums2 {
		if _, ok := map_2[v]; !ok {
			map_2[v] = empty
		}
	}
	var res_1 []int
	for k := range map_1 {
		if _, ok := map_2[k]; !ok {
			res_1 = append(res_1, k)
		}
	}
	var res_2 []int

	for k := range map_2 {
		if _, ok := map_1[k]; !ok {
			res_2 = append(res_2, k)
		}
	}

	res = append(res, res_1, res_2)
	return res
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	findDifference(nums1, nums2)
}
