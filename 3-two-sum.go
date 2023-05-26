package main

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func twoSum(nums []int, target int) []int {
	// m := make(map[int]int)
	// for i, v := range nums {
	// 	if j, ok := m[target-v]; ok {
	// 		return []int{j, i}
	// 	}
	// 	m[v] = i
	// }
	// return nil

	var m *orderBTree
	for i, v := range nums {
		if find := m.find(target - v); find != nil {
			return []int{find.order, i}
		}
		m = m.insert(i, v)
	}
	return nil
}

type orderBTree struct {
	order int
	//	count       int
	n           int
	left, right *orderBTree
}

func (m *orderBTree) insert(order, n int) *orderBTree {
	if m == nil {
		m = new(orderBTree)
		m.order = order
		m.n = n
		return m
	}
	if n == m.n {
		m.order = order
		return m
	}
	if n < m.n {
		m.left = m.left.insert(order, n)
		return m
	}
	m.right = m.right.insert(order, n)
	return m
}

func (m *orderBTree) find(n int) *orderBTree {
	if m == nil {
		return nil
	}
	if n == m.n {
		return m
	}
	if n < m.n {
		return m.left.find(n)
	}
	return m.right.find(n)
}
