package main

import (
	"fmt"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// outputs := [][]string{
	// 	{"bat"},
	// 	{"nat", "tan"},
	// 	{"ate", "eat", "tea"},
	// }
	fmt.Println(groupAnagrams(strs))

	strs1 := []string{"a"}
	// outputs1 := [][]string{
	// 	{"a"},
	// }
	fmt.Println(groupAnagrams(strs1))

	strs2 := []string{""}
	// outputs2 := [][]string{
	// 	{""},
	// }
	fmt.Println(strs2)
}
