package user

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name          string
		inputName     string
		inputEmail    string
		expectedName  string
		expectedEmail string
	}{
		{"normal user", "John Doe", "John@Example.com", "John Doe", "john@example.com"},
		{"whitespace trim", " Jane Smith ", " JANE@EXAMPLE.COM ", "Jane Smith", "jane@example.com"},
		{"empty name", "", "test@example.com", "", "test@example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := New(tt.inputName, tt.inputEmail)

			if user.Name != tt.expectedName {
				t.Errorf("New() name = %q; want %q", user.Name, tt.expectedName)
			}

			if user.Email != tt.expectedEmail {
				t.Errorf("New() email = %q; want %q", user.Email, tt.expectedEmail)
			}
		})
	}
}

func TestUser_IsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"valid email", "test@example.com", true},
		{"valid email with subdomain", "user@mail.example.com", true},
		{"valid email with numbers", "user123@example123.com", true},
		{"valid email with special chars", "user.name+tag@example.com", true},
		{"invalid - no @", "testexample.com", false},
		{"invalid - no domain", "test@", false},
		{"invalid - no local part", "@example.com", false},
		{"invalid - no TLD", "test@example", false},
		{"invalid - spaces", "test @example.com", false},
		{"empty email", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{Email: tt.email}
			result := user.IsValidEmail()

			if result != tt.expected {
				t.Errorf("IsValidEmail() for %q = %t; want %t", tt.email, result, tt.expected)
			}
		})
	}
}

func TestUser_GetDisplayName(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		expected string
	}{
		{"normal name", "John Doe", "John Doe"},
		{"empty name", "", "Anonymous"},
		{"whitespace only", "   ", "   "}, // trimming is done in New(), not GetDisplayName()
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{Name: tt.userName}
			result := user.GetDisplayName()

			if result != tt.expected {
				t.Errorf("GetDisplayName() = %q; want %q", result, tt.expected)
			}
		})
	}
}

func TestUserIntegration(t *testing.T) {
	user := New("John Doe", "john@example.com")

	if !user.IsValidEmail() {
		t.Error("Expected valid email for john@example.com")
	}

	if user.GetDisplayName() != "John Doe" {
		t.Errorf("Expected display name 'John Doe', got %q", user.GetDisplayName())
	}
}
