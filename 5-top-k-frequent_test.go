package main

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums1 := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums1, 200))
	// outputs: [1, 2]

	nums2 := []int{1}
	fmt.Println(topKFrequent(nums2, 1))
	// outputs: [1]
}
