package main

/*
	Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
	cannot use additional data structures?
*/

func HasAllUniqueCharacters(word string) bool {
	var s = make(map[rune]int)
	for _, char := range word {
		if s[char] != 0 {
			return false
		}
		s[char] = 1
	}
	return true
}
