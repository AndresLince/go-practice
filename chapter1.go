package main

import (
	"strings"
)

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

/*
	URLify: Write a method to replace all spaces in a string with '%20'. You
	may assume that the string has sufficient space at the end to hold the
	additional characters,and that you are given the "true" length of the
	string. (Note: If implementing in Java, please use a character array so
	that you can perform this operation in place.)
*/
func Urlify(word string, length int) string {
	emptyChars := countChars(word, 32, 0, length)
	longIndex := length + (emptyChars * 2) - 1
	shortIndex := length - 1
	for shortIndex >= 0 {
		if word[shortIndex] == 32 {
			word = setCharAt(word, longIndex, 37)
			word = setCharAt(word, longIndex-1, 48)
			word = setCharAt(word, longIndex-2, 50)
			longIndex = longIndex - 3
		} else {
			word = setCharAt(word, longIndex, rune(word[shortIndex]))
			longIndex--
		}
		shortIndex--
	}
	return word
}

func countChars(word string, char byte, from int, to int) int {
	counter := 0
	for from < to {
		if char == word[from] {
			counter++
		}
		from++
	}
	return counter
}

func setCharAt(word string, index int, char rune) string {
	if index > len(word) {
		return word
	}
	return word[:index] + string(char) + word[index+1:]
}

/*
	Given a string, write a function to check if it is a permutation of a
	palindrome. A palindrome is a word or phrase that is the same forwards and
	backwards. A permutation is a rearrangement of letters. The palindrome does
	not need to be limited to just dictionary words.
*/
func IsAPalindromePermutation(palindrome string) bool {
	//var letters = make(map[rune]int)
	var letters = make(map[rune]int)
	palindrome = strings.ToLower(palindrome)
	for _, char := range palindrome {
		if char == 32 {
			continue
		}
		if letters[char] == 0 {
			letters[char] = 1
		} else {
			letters[char]++
		}
	}
	oddCounter := 0
	for _, char := range letters {
		if isOddNumber(char) {
			if oddCounter == 1 {
				return false
			}
			oddCounter++
		}
	}
	return true
}

func isOddNumber(number int) bool {
	return number%2 != 0
}
