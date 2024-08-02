package main

import (
	"regexp"
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
        {"LongPassword", "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ", 33},
        {"MissingLowercase", "AA11", 2},
        {"MissingUppercase", "aa11", 2},
        {"MissingDigit", "aaAA", 2},
        {"RepeatingCharacters", "aaa111", 1},
        {"ComplexRepeatingCharacters", "aaa111AAA", 2},
        {"JustRightStrongPassword", "aA1aA1aA1", 0},
        {"AlmostStrongPassword", "aaaaA1", 1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := calculateSteps(tt.password); got != tt.want {
                t.Errorf("calculateSteps() = %v, want %v", got, tt.want)
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
        name          string
        password      string
        wantLower     bool
        wantUpper     bool
        wantDigit     bool
        wantNoRepeat  bool
    }{
        {"AllCriteriaMet", "aA1bB2cC3", true, true, true, true},
        {"MissingLowercase", "AA11BB22", false, true, true, true},
        {"MissingUppercase", "aa11bb22", true, false, true, true},
        {"MissingDigit", "aAbbCCdd", true, true, false, true},
        {"RepeatingCharacters", "aA111bB", true, true, true, false},
        {"OnlyLowercase", "abcdefg", true, false, false, true},
        {"OnlyUppercase", "ABCDEFG", false, true, false, true},
        {"OnlyDigits", "1234567", false, false, true, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            hasLower := false
            hasUpper := false
            hasDigit := false
            for _, char := range tt.password {
                if char >= 'a' && char <= 'z' {
                    hasLower = true
                } else if char >= 'A' && char <= 'Z' {
                    hasUpper = true
                } else if char >= '0' && char <= '9' {
                    hasDigit = true
                }
            }
            noRepeat := !regexp.MustCompile(`(.)\1\1`).MatchString(tt.password)

            if hasLower != tt.wantLower {
                t.Errorf("Lowercase check for %q = %v, want %v", tt.password, hasLower, tt.wantLower)
            }
            if hasUpper != tt.wantUpper {
                t.Errorf("Uppercase check for %q = %v, want %v", tt.password, hasUpper, tt.wantUpper)
            }
            if hasDigit != tt.wantDigit {
                t.Errorf("Digit check for %q = %v, want %v", tt.password, hasDigit, tt.wantDigit)
            }
            if noRepeat != tt.wantNoRepeat {
                t.Errorf("No repeat check for %q = %v, want %v", tt.password, noRepeat, tt.wantNoRepeat)
            }
        })
    }
}