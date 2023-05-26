package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	duplicates := [][]int{
		{7, 9, 100, 22, 30000, 1, 2, 3, 1},
		{19, 12, 249, 8, 30, 1, 2, 3, 4, 3},
		{39, 23, 100, 470, 470, 1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}
	uniques := [][]int{
		{1, 2, 100, 97, 32, 38, 49, 3, 4},
		{1, 2, 700, 197, 3332, 98, 69, 3, 4, 5},
		{},
	}
	for _, v := range duplicates {
		if !containsDuplicate(v) {
			t.Errorf("%v Expected duplicate/true, got unique/false", v)
		}
	}
	for _, v := range uniques {
		if containsDuplicate(v) {
			t.Errorf("%v Expected unique/false, got duplicate/true", v)
		}
	}
}
