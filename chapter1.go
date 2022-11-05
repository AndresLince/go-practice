package main

/*
	Is Unique: Implement an algorithm to determine if a string has all unique
	characters. What if you
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

/*
	Check Permutation: Given two strings, write a method to decide if one is a
	permutation of the other.
*/

func CheckPermutation(word1, word2 string) bool {
	string1Length := len([]rune(word1))
	string2Length := len([]rune(word2))
	if string1Length != string2Length {
		return false
	}
	var letters = make(map[rune]int)
	for _, char := range word1 {
		if letters[char] != 0 {
			return false
		}
		letters[char] = 1
	}
	for _, char2 := range word2 {
		if letters[char2] == 0 {
			return false
		}
		letters[char2]--
		if letters[char2] < 0 {
			return false
		}
	}
	return true
}
