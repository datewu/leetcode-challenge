package main

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	buckets := make([][]int, len(nums)+1)
	for mum, count := range m {
		buckets[count] = append(buckets[count], mum)
	}
	res := make([]int, 0)
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		if buckets[i] == nil {
			continue
		}
		res = append(res, buckets[i]...)
	}
	// return res[:k]
	return res
}
