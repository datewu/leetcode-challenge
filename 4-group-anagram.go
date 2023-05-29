package main

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once
// strs[i] consists of lowercase English letters.
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[int][]string)
	for _, str := range strs {
		key := getAnagramstKey(str)
		m[key] = append(m[key], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// good enough, to be proved
func getAnagramstKey(str string) int {
	m := make(map[rune]int)
	for _, r := range str {
		m[r]++
	}
	var key int
	// [a, b]
	for k, v := range m {
		key += v + int(k) + 1
	}
	// [a-c, a+c] => c = a-b => [b, a]
	for k, v := range m {
		key *= v * (int(k) + 1)
	}
	return key + 100_000*len(str) + 1000_000*len(m)
}
