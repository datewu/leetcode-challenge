package main

import "testing"

func TestValidAnagram(t *testing.T) {
	sucesses := [][2]string{
		{"anagram", "nagaram"},
		{"iam your father!", "father iam your!"},
		{"iam your father!", "father your ami!"},
		{"我是你爸爸！👨", "👨我是你！爸爸"},
		{"a", "a"},
		{"ab", "ba"},
		{"ab", "ab"},
	}

	failds := [][2]string{
		{"anagram", "nagaramm"},
		{"rat", "car"},
		{"rat", "carrr"},
		{"a", "aa"},
		{"a", "A"}, // case matters
		{"iam your father!", "father iam you !"},
		{"ab", "bb"},
		{"ab", "a"},
		{"我是你爸爸。👨", "👨我是你！爸爸"},
	}

	for _, v := range sucesses {
		if !validAnagramWithTree(v[0], v[1]) {
			t.Errorf("Tree: %v Expected true, got false", v)
		}
		if !validAnagramWithKey(v[0], v[1]) {
			t.Errorf("Key: %v Expected true, got false", v)
		}
	}
	for _, v := range failds {
		if validAnagramWithTree(v[0], v[1]) {
			t.Errorf("Tree: %v Expected false, got true", v)
		}
		if validAnagramWithKey(v[0], v[1]) {
			t.Errorf("Key: %v Expected true, got false", v)
		}
	}
}
