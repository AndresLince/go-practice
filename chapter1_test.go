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
