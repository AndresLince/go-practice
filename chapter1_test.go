package main

import (
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
