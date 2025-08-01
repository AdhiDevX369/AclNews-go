package calculator

import "fmt"

type Calculator struct{}

func New() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
