package cmd

import (
	"strings"
	"testing"
)

func TestGeneratePasswordDefaultIsAlphaNumeric(t *testing.T) {
	password, err := generatePassword(16, false)
	if err != nil {
		t.Fatalf("generatePassword returned error: %v", err)
	}

	if len(password) != 16 {
		t.Fatalf("expected password length 16, got %d", len(password))
	}

	allowed := passwdLowerCharset + passwdUpperCharset + passwdDigitCharset
	for _, ch := range password {
		if !strings.ContainsRune(allowed, ch) {
			t.Fatalf("password contains unexpected character %q", ch)
		}
	}
}

func TestGeneratePasswordComplexContainsAllClasses(t *testing.T) {
	password, err := generatePassword(20, true)
	if err != nil {
		t.Fatalf("generatePassword returned error: %v", err)
	}

	if len(password) != 20 {
		t.Fatalf("expected password length 20, got %d", len(password))
	}

	if !containsAny(password, passwdLowerCharset) {
		t.Fatal("expected complex password to contain a lowercase character")
	}
	if !containsAny(password, passwdUpperCharset) {
		t.Fatal("expected complex password to contain an uppercase character")
	}
	if !containsAny(password, passwdDigitCharset) {
		t.Fatal("expected complex password to contain a digit")
	}
	if !strings.ContainsRune(passwdSymbolCharset, rune(password[len(password)-1])) {
		t.Fatal("expected complex password to end with a symbol")
	}
	if countCharsFromSet(password, passwdSymbolCharset) != 1 {
		t.Fatal("expected complex password to contain exactly one symbol")
	}
}

func TestGeneratePasswordRejectsInvalidLength(t *testing.T) {
	if _, err := generatePassword(0, false); err == nil {
		t.Fatal("expected error for non-positive length")
	}

	if _, err := generatePassword(3, true); err == nil {
		t.Fatal("expected error for complex password shorter than required classes")
	}

	password, err := generatePassword(4, true)
	if err != nil {
		t.Fatalf("expected minimum complex password length to be valid: %v", err)
	}
	if len(password) != 4 {
		t.Fatalf("expected password length 4, got %d", len(password))
	}
}

func containsAny(value string, charset string) bool {
	for _, ch := range value {
		if strings.ContainsRune(charset, ch) {
			return true
		}
	}
	return false
}

func countCharsFromSet(value string, charset string) int {
	count := 0
	for _, ch := range value {
		if strings.ContainsRune(charset, ch) {
			count++
		}
	}
	return count
}
