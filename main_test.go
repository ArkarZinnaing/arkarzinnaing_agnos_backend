package main

import (
	"testing"
)

func TestCalculateSteps(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     int
	}{
		{"Example1", "aA1", 3},
		{"Example2", "1445D1cd", 0},
		{"ShortPassword", "a", 5},
		{"LongPassword", "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ", 34},
		{"MissingLowercase", "AA11", 2},
		{"MissingUppercase", "aa11", 2},
		{"MissingDigit", "aaAA", 2},
		{"RepeatingCharacters", "aaa111", 1},
		{"ComplexRepeatingCharacters", "aaa111AAA", 2},
		{"JustRightStrongPassword", "aA1aA1", 0},
		{"AlmostStrongPassword", "aaaaA1", 1},
		{"MaxLengthPassword", "aA1aA1aA1aA1aA1aA1aA", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSteps(tt.password); got != tt.want {
				t.Errorf("calculateSteps(%q) = %v, want %v", tt.password, got, tt.want)
			}
		})
	}
}

func TestPasswordLengthCriteria(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"TooShort", "aA1", false},
		{"MinimumLength", "aA1aA1", true},
		{"MaximumLength", "aA1aA1aA1aA1aA1aA1aA", true},
		{"TooLong", "aA1aA1aA1aA1aA1aA1aA1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := len(tt.password) >= 6 && len(tt.password) <= 20
			if got != tt.want {
				t.Errorf("Length check for %q = %v, want %v", tt.password, got, tt.want)
			}
		})
	}
}

func TestPasswordCharacterCriteria(t *testing.T) {
	tests := []struct {
		name      string
		password  string
		wantLower bool
		wantUpper bool
		wantDigit bool
	}{
		{"AllCriteriaMet", "aA1bB2", true, true, true},
		{"MissingLowercase", "AA11BB", false, true, true},
		{"MissingUppercase", "aa11bb", true, false, true},
		{"MissingDigit", "aAbbCC", true, true, false},
		{"OnlyLowercase", "abcdef", true, false, false},
		{"OnlyUppercase", "ABCDEF", false, true, false},
		{"OnlyDigits", "123456", false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasLower, hasUpper, hasDigit := checkCharacterTypes(tt.password)
			if hasLower != tt.wantLower {
				t.Errorf("Lowercase check for %q = %v, want %v", tt.password, hasLower, tt.wantLower)
			}
			if hasUpper != tt.wantUpper {
				t.Errorf("Uppercase check for %q = %v, want %v", tt.password, hasUpper, tt.wantUpper)
			}
			if hasDigit != tt.wantDigit {
				t.Errorf("Digit check for %q = %v, want %v", tt.password, hasDigit, tt.wantDigit)
			}
		})
	}
}

func TestRepeatingCharacters(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{"NoRepeats", "abcdef", false},
		{"TwoRepeats", "aabbcc", false},
		{"ThreeRepeats", "aaa123", true},
		{"MultipleThreeRepeats", "aaa111bbb", true},
		{"MixedRepeats", "aa111b", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasRepeatingCharacters(tt.password)
			if got != tt.want {
				t.Errorf("hasRepeatingCharacters(%q) = %v, want %v", tt.password, got, tt.want)
			}
		})
	}
}

func checkCharacterTypes(password string) (bool, bool, bool) {
	hasLower, hasUpper, hasDigit := false, false, false
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			hasLower = true
		} else if char >= 'A' && char <= 'Z' {
			hasUpper = true
		} else if char >= '0' && char <= '9' {
			hasDigit = true
		}
	}
	return hasLower, hasUpper, hasDigit
}

func hasRepeatingCharacters(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i] == password[i+2] {
			return true
		}
	}
	return false
}
