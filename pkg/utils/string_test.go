package utils

import "testing"

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"normal string", "hello", "Hello"},
		{"already capitalized", "Hello", "Hello"},
		{"single character", "a", "A"},
		{"empty string", "", ""},
		{"unicode", "ñoño", "Ñoño"},
		{"with spaces", "hello world", "Hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Capitalize(tt.input)
			if result != tt.expected {
				t.Errorf("Capitalize(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"normal string", "hello", "olleh"},
		{"palindrome", "racecar", "racecar"},
		{"single character", "a", "a"},
		{"empty string", "", ""},
		{"with spaces", "hello world", "dlrow olleh"},
		{"unicode", "café", "éfac"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"simple palindrome", "racecar", true},
		{"not palindrome", "hello", false},
		{"palindrome with spaces", "race car", true},
		{"palindrome mixed case", "RaceCar", true},
		{"single character", "a", true},
		{"empty string", "", true},
		{"phrase palindrome", "A man a plan a canal Panama", true}, // this is actually a palindrome when spaces are removed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %t; want %t", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkCapitalize(b *testing.B) {
	text := "hello world this is a benchmark test"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Capitalize(text)
	}
}

func BenchmarkReverse(b *testing.B) {
	text := "hello world this is a benchmark test"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse(text)
	}
}
