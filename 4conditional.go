// This program demonstrates all types of conditional statements in Go.

package main

import "fmt"

func main() {

	// 1. Simple if statement
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	// 2. if-else statement
	age2 := 15
	if age2 >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}


	// 3. if-else if-else chain
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// 4. Short if statement (variable declared inside the condition)
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 5. Nested if statements
	age3 := 25
	citizen := true
	if age3 >= 18 {
		if citizen {
			fmt.Println("You can vote.")
		} else {
			fmt.Println("You must be a citizen to vote.")
		}
	} else {
		fmt.Println("You are not old enough to vote.")
	}

	// 6. Switch statement (with string cases)
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend loading...")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Midweek grind")
	}

	// 7. Switch without an expression (acts like if-else chain)
	num := 42
	switch {
	case num < 0:
		fmt.Println("Negative number")
	case num == 0:
		fmt.Println("Zero")
	default:
		fmt.Println("Positive number")
	}

	// 8. Switch with fallthrough (manually continues to the next case)
	number := 1
	switch number {
	case 1:
		fmt.Println("One")
		fallthrough // allows the next case to run even if matched
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Something else")
	}
}
