package calculator

import (
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"zero", 0, 5, 5},
		{"mixed", -5, 10, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 10, 3, 7},
		{"negative numbers", -5, -3, -2},
		{"zero", 5, 0, 5},
		{"negative result", 3, 10, -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 4, 3, 12},
		{"negative numbers", -4, -3, 12},
		{"zero", 0, 5, 0},
		{"mixed", -4, 3, -12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		a, b        int
		expected    int
		expectError bool
	}{
		{"positive numbers", 12, 3, 4, false},
		{"negative numbers", -12, -3, 4, false},
		{"mixed", -12, 3, -4, false},
		{"division by zero", 10, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("Divide(%d, %d) expected error but got none", tt.a, tt.b)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

func BenchmarkCalculator_Add(b *testing.B) {
	calc := New()
	for i := 0; i < b.N; i++ {
		calc.Add(100, 200)
	}
}
