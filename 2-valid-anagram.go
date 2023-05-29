package main

type countBTree struct {
	n, count    int
	left, right *countBTree
}

func (t *countBTree) insert(n int) *countBTree {
	if t == nil {
		t = new(countBTree)
		t.n = n
		t.count = 1
		return t
	}
	if n == t.n {
		t.count++
		return t
	}
	if n < t.n {
		t.left = t.left.insert(n)
		return t
	}
	t.right = t.right.insert(n)
	return t
}

// find n and decrement count, which mutate the origin countBtree
// use with caution
func (t *countBTree) find(n int) bool {
	if t == nil {
		return false
	}
	if n == t.n {
		if t.count > 0 {
			t.count--
			return true
		}
		return false
	}
	if n < t.n {
		return t.left.find(n)
	}
	return t.right.find(n)
}

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
func validAnagramWithTree(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var tree *countBTree
	for _, v := range s {
		tree = tree.insert(int(v))
	}
	for _, v := range t {
		if !tree.find(int(v)) {
			return false
		}
	}
	return true
}

func validAnagramWithKey(s string, t string) bool {
	// if len(s) != len(t) {
	// 	return false
	// }
	return getAnagramstKey(s) == getAnagramstKey(t)
}
