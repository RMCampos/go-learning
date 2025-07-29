package main

import (
	"fmt"
	"strings"
)

// Exercise 1: FizzBuzz
// Print numbers 1-100, but replace multiples of 3 with "Fizz"
// multiple of 5 with "Buzz", and multiple of both with "FizzBuzz"
func fizzBuzz() {
	fmt.Println("== Exercise 1: FizzBuzz ==")
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("FizzBuzz ")
		case i%3 == 0:
			fmt.Printf("Fizz ")
		case i%5 == 0:
			fmt.Printf("Buzz ")
		default:
			fmt.Printf("%d", i)
		}
		if i%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// Exercise 2: Palindrome Checker
// Check if a string reads the same forwards and backwards
func isPalindrome(s string) bool {
	// Convert to lowercase and remove spaces for easier comparision
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	// Convert to runes to properly handle Unicode characters
	runes := []rune(s)

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

// Exercise 3: Factorial Calulator
// Calculate the factorial of a number (n!)
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Exercise 4: Find Maximum in Slice
// Find the largest number in a slice of integers
func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// Exercise 5: Count Vowels
// Count the number of vowels in a string
func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0

	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	// Test all exercises
	fmt.Printf("Go Programming Practice Exercises\n\n")

	// Exercise 1
	fizzBuzz()

	// Exercise 2
	fmt.Println("== Exercise 2: Palindrome Checker ==")
	testStrings := []string{"racecar", "hello", "A man a plan a canal Panama", "race a car"}
	for _, s := range testStrings {
		fmt.Printf("'%s' is palindrome: %t\n", s, isPalindrome(s))
	}
	fmt.Println()

	// Exercise 3
	fmt.Println("== Exercise 3: Factorial ==")
	for i := 0; i <= 6; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
	fmt.Println()

	// Exercise 4
	fmt.Println("== Exercise 4: Find Maximum ==")
	numbers := []int{27, 3, 7, 2, 9, 1, 8, 5}
	fmt.Printf("Max in %v: %d\n", numbers, findMax(numbers))
	fmt.Println()

	// Exercise 5
	fmt.Println("== Exercise 5: Count Vowels ==")
	text := "My name is Ricardo Campos"
	fmt.Printf("Vowels in '%s': %d\n", text, countVowels(text))
	fmt.Println()
}
