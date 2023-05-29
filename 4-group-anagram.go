package main

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once
// strs[i] consists of lowercase English letters.
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[[26]int][]string)
	for _, str := range strs {
		key := getAnagramstKey(str)
		m[key] = append(m[key], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func getAnagramstKey(str string) [26]int {
	m := [26]int{}

	for _, r := range str {
		m[r%26]++
	}
	return m
}
