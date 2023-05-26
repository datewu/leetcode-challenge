package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	sucesses := map[int][]int{
		9:     {1, 1, 1, 1, 1, 3, 3, 33, 6, 18, 9},
		23:    {20, 1, 2, 2, 2, 2, 29, 3, 887},
		145:   {200, 89, 23, 35, 207, 200, 9, 23, 13, 34, 200, 205, 201, 0, 145},
		236:   {200, 89, 23, 35, 200, 9, 1, 23, 223, 812, 89, 0, 36, 36, 36, 199},
		89843: {89, 234, 898, 89840, 33, 7, 3, 9},
	}
	failds := map[int][]int{
		9:     {1, 1, 1, 1, 1, 18, 9, 98},
		23:    {28, 1, 2, 2, 2, 2, 29, 3, 887},
		235:   {200, 89, 23, 36, 200, 200, 9},
		89843: {89, 234, 898, 89841, 33, 7, 3, 9},
		0:     {0, 1, 2, 3, 4, 4, 5534, 14},
	}
	for k, v := range sucesses {
		idxs := twoSum(v, k)
		if idxs == nil {
			t.Errorf("%d => %v Expected true, got false", k, v)
			continue
		}
		t.Logf("%v : %v => %d \n%d + %d = %d", v, idxs, k, v[idxs[0]], v[idxs[1]], k)
	}
	for k, v := range failds {
		idxs := twoSum(v, k)
		if idxs != nil {
			t.Logf("%d + %d = %d", v[idxs[0]], v[idxs[1]], k)
			t.Errorf("%d => %v Expected false, got true", k, v)
		}
	}
}
