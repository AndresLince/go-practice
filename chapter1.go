package main

import (
	"math"
	"strconv"
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

	if string1Length == 0 {
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

/*
	There are three types of edits that can be performed on strings: insert a
	character, remove a character, or replace a character. Given two strings,
	write a function to check if they are one edit (or zero edits) away.
*/
func OneEditAway(word1, word2 string) bool {
	length1 := len(word1)
	length2 := len(word2)
	if math.Abs(float64(length1-length2)) > 1 {
		return false
	}
	if length1 == length2 {
		return validateOneEdit(word1, word2)
	}
	if length1 < length2 {
		return validateOneInsert(word1, word2)
	}

	return validateOneInsert(word2, word1)
}

func validateOneEdit(word1, word2 string) bool {
	diffCounter := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			if diffCounter == 1 {
				return false
			}
			diffCounter++
		}
	}
	return true
}

func validateOneInsert(word1, word2 string) bool {
	index1 := 0
	index2 := 0
	for index1 < len(word1) && index2 < len(word2) {
		if word1[index1] != word2[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}

	return true
}

/*
	String Compression: Implement a method to perform basic string compression
	using the counts of repeated characters. For example, the string
	aabcccccaaa would become a2blc5a3. If the "compressed" string would not
	become smaller than the original string, your method should return the
	original string. You can assume the string has only uppercase and lowercase
	letters (a - z).
*/

func StringCompression(word string) string {
	newString := ""
	lastChar := word[0]
	counter := 0
	for _, char := range word {
		if byte(char) != lastChar {
			newString = newString + string(lastChar) + strconv.Itoa(counter)
			counter = 0
		}
		lastChar = byte(char)
		counter++
	}
	newString = newString + string(lastChar) + strconv.Itoa(counter)
	if len(newString) < len(word) {
		return newString
	}

	return word
}

/*
	Rotate Matrix: Given an image represented by an NxN matrix, where each
	pixel in the image is 4 bytes, write a method to rotate the image by 90
	degrees. Can you do this in place?
*/
func RotateMatrix(matrix Nums) Nums {
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]
			matrix[first][i] = matrix[last-offset][first]
			matrix[last-offset][first] = matrix[last][last-offset]
			matrix[last][last-offset] = matrix[i][last]
			matrix[i][last] = top
		}
	}
	return matrix
}
