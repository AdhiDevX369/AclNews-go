package user

import (
	"regexp"
	"strings"
)

type User struct {
	Name  string
	Email string
}

func New(name, email string) *User {
	return &User{
		Name:  strings.TrimSpace(name),
		Email: strings.ToLower(strings.TrimSpace(email)),
	}
}

func (u *User) IsValidEmail() bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(u.Email)
}

func (u *User) GetDisplayName() string {
	if u.Name == "" {
		return "Anonymous"
	}
	return u.Name
}
