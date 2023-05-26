package main

import "fmt"

type bTree struct {
	n           int
	left, right *bTree
}

func (t *bTree) walk() []int {
	if t == nil {
		return nil
	}
	return append(append(t.left.walk(), t.n), t.right.walk()...)
}

func (t *bTree) insert(n int) (*bTree, bool) {
	if t == nil {
		t = new(bTree)
		t.n = n
		return t, true
	}
	if n == t.n {
		return t, false
	}
	if n < t.n {
		left, inserted := t.left.insert(n)
		t.left = left
		return t, inserted
	}
	right, inserted := t.right.insert(n)
	t.right = right
	return t, inserted
}

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
func containsDuplicate(nums []int) bool {
	var t *bTree
	for _, v := range nums {
		n, inserted := t.insert(v)
		if !inserted {
			// fmt.Println(t.walk())
			return true
		}
		t = n
	}
	fmt.Println(t.walk())
	return false
}
