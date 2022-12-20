package main

import (
	"reflect"
	"testing"
)

// TestHasAllUniqueCharacters calls chapter1.HasAllUniqueCharacters with a
// word, checking for a valid return value.
func TestHasAllUniqueCharacters(t *testing.T) {
	response := HasAllUniqueCharacters("algo")
	if response != true {
		t.Fatalf(`HasAllUniqueCharacters("algo") expect true`)
	}
	response = HasAllUniqueCharacters("algoa")
	if response != false {
		t.Fatalf(`HasAllUniqueCharacters("algoa") expect false`)
	}
}
func TestCheckPermutation(t *testing.T) {
	response := CheckPermutation("algo", "loga")
	if response != true {
		t.Fatalf(`CheckPermutation("algo", "loga") expect true`)
	}
	response = CheckPermutation("algoa", "algo")
	if response != false {
		t.Fatalf(`CheckPermutation("algoa", "algo") expect false`)
	}
	response = CheckPermutation("", "")
	if response != false {
		t.Fatalf(`CheckPermutation("", "") expect false`)
	}
	response = CheckPermutation("Dabale arroz a la zorra el abad", "Dabale arroz a la zorra el abad")
	if response != false {
		t.Fatalf(`CheckPermutation("", "") expect false`)
	}
	response = CheckPermutation("avc", "ase")
	if response != false {
		t.Fatalf(`CheckPermutation("abcdef", "abcdef") expect false`)
	}
}

func TestUrlify(t *testing.T) {
	response := Urlify("Mr John Smith    ", 13)
	if response != "Mr20%John20%Smith" {
		t.Fatalf(`Urlify("Mr John Smith    ", 13) expect valid string`)
	}
}

func TestIsAPalindromePermutation(t *testing.T) {
	response := IsAPalindromePermutation("Tact Coa")
	if response != true {
		t.Fatalf(`IsAPalindromePermutation("Tact Coa") expect true`)
	}
	response = IsAPalindromePermutation("taco cat")
	if response != true {
		t.Fatalf(`IsAPalindromePermutation("taco cat") expect true`)
	}
	response = IsAPalindromePermutation("taco cat")
	if response != true {
		t.Fatalf(`IsAPalindromePermutation("atco cta") expect true`)
	}
	response = IsAPalindromePermutation("Dabale arroz a la zorra el abad")
	if response != true {
		t.Fatalf(`IsAPalindromePermutation("Dabale arroz a la zorra el abad") expect true`)
	}
	response = IsAPalindromePermutation("Dabale carne a la zorra el abad")
	if response != false {
		t.Fatalf(`IsAPalindromePermutation("Dabale carne a la zorra el abad") expect false`)
	}
}

func TestOneEditAway(t *testing.T) {
	response := OneEditAway("pale", "ple")
	if response != true {
		t.Fatalf(`OneEditAway("pale", "ple") expect true`)
	}
	response = OneEditAway("pales", "pale")
	if response != true {
		t.Fatalf(`OneEditAway("pales", "pale") expect true`)
	}
	response = OneEditAway("pale", "bale")
	if response != true {
		t.Fatalf(`OneEditAway("pale", "bale") expect true`)
	}
	response = OneEditAway("pale", "bake")
	if response != false {
		t.Fatalf(`OneEditAway("pale", "bake") expect true`)
	}
	response = OneEditAway("apple", "aple")
	if response != true {
		t.Fatalf(`OneEditAway("apple", "aple") expect true`)
	}
	response = OneEditAway("aple", "apple")
	if response != true {
		t.Fatalf(`OneEditAway("aple", "apple") expect true`)
	}
	response = OneEditAway("aplese", "apple")
	if response != false {
		t.Fatalf(`OneEditAway("aplese", "apple") expect true`)
	}
	response = OneEditAway("apleses", "apple")
	if response != false {
		t.Fatalf(`OneEditAway("apleses", "apple") expect true`)
	}
	response = OneEditAway("apple", "apleses")
	if response != false {
		t.Fatalf(`OneEditAway("apple", "apleses") expect true`)
	}
}

func TestStringCompression(t *testing.T) {
	response := StringCompression("aabccccaaa")
	if response != "a2b1c4a3" {
		t.Fatalf(`StringCompression("aabccccaaa") expect "a2b1c4a3"`)
	}
	response = StringCompression("aabbabbccccaaa")
	if response != "a2b2a1b2c4a3" {
		t.Fatalf(`StringCompression("aabbabbccccaaa") expect "a2b2a1b2c4a3"`)
	}
	response = StringCompression("aabbabbccccaaacdeal")
	if response != "aabbabbccccaaacdeal" {
		t.Fatalf(`StringCompression("aabbabbccccaaacdeal") expect "aabbabbccccaaacdeal"`)
	}
}

func TestRotateMatriz(t *testing.T) {
	matrix := Nums{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}
	correctResult := Nums{
		[]int{13, 9, 5, 1},
		[]int{14, 10, 6, 2},
		[]int{15, 11, 7, 3},
		[]int{16, 12, 8, 4},
	}

	response := RotateMatrix(matrix)
	if !reflect.DeepEqual(response, correctResult) {
		t.Fatalf(`RotateMatrix(matrix) expect a valid matriz`)
	}
	matrix = Nums{
		[]int{1, 2, 3, 4, 5},
		[]int{6, 7, 8, 9, 10},
		[]int{11, 12, 13, 14, 15},
		[]int{16, 17, 18, 19, 20},
		[]int{21, 22, 23, 24, 25},
	}
	correctResult = Nums{
		[]int{21, 16, 11, 6, 1},
		[]int{22, 17, 12, 7, 2},
		[]int{23, 18, 13, 8, 3},
		[]int{24, 19, 14, 9, 4},
		[]int{25, 20, 15, 10, 5},
	}
	response = RotateMatrix(matrix)
	if !reflect.DeepEqual(response, correctResult) {
		t.Fatalf(`RotateMatrix(matrix) expect a valid matriz`)
	}
}
