package main

import "fmt"

// TODO: Complete the following exercises to practice Go!

func main() {
	fmt.Println("🎯 Go Practice Exercises")
	fmt.Println("========================")

	// Exercise 1: Create a greeting function
	// TODO: Uncomment and complete this function call
	greeting := createGreeting("Your Name")
	fmt.Println(greeting)

	// Exercise 2: Work with numbers
	// TODO: Create variables for your age and birth year
	// Calculate and print what year you'll turn 100

	// Exercise 3: Favorite things
	// TODO: Create a slice of your 5 favorite movies
	// Print each movie with its position (1. Movie1, 2. Movie2, etc.)

	// Exercise 4: Grade calculator
	// TODO: Create a map of subjects and grades
	// Calculate and print the average grade

	// Exercise 5: Simple calculator
	// TODO: Use the calculator functions below
}

// Exercise 1: Complete this function
// TODO: Make this function return a personalized greeting
func createGreeting(name string) string {
	// Return something like "Hello, [name]! Welcome to Go!"
	return ""
}

// Exercise 2: Create these calculator functions
// TODO: Implement these functions

func multiply(a, b int) int {
	// TODO: Return a * b
	return 0
}

func subtract(a, b int) int {
	// TODO: Return a - b
	return 0
}

// Exercise 3: Working with slices
// TODO: Create a function that takes a slice of strings and returns the longest one
func findLongest(words []string) string {
	// TODO: Find and return the longest word
	return ""
}

// Exercise 4: Working with maps
// TODO: Create a function that calculates average from a map of grades
func calculateAverage(grades map[string]float64) float64 {
	// TODO: Calculate the average of all grades
	return 0.0
}

// BONUS Exercise: Create a simple guessing game!
// TODO: Think of a number, ask user input, tell if too high/low
// You'll need to import "bufio", "os", "strconv" for user input
