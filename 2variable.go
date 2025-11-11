// In this file, you will learn how to declare variables and constants in Go (Golang)

package main

import "fmt"

func main() {
	// VARIABLES

	// 1. Declare variable with explicit type
	var name1 string = "Hola Amigo"
	fmt.Println(name1)

	// 2. Declare variable without specifying type (Go infers it)
	var name2 = "Arigato"
	fmt.Println(name2)

	// 3. Shorthand declaration (only works inside functions)
	name3 := "Namaste"
	fmt.Println(name3)

	// 4. Declare multiple variables in one line
	var firstName, lastName = "Arpit", "Sarang"
	fmt.Println(firstName, lastName)

	// 5. Declare multiple variables in a block
	var (
		country = "India"
		age     = 20
		coder   = true
	)
	fmt.Println(country, age, coder)

	// CONSTANTS

	// 6. Single constant
	const pi = 3.14159
	fmt.Println("Value of Pi:", pi)

	// 7. Constant with explicit type
	const gravity float64 = 9.8
	fmt.Println("Gravity:", gravity)

	// 8. Multiple constants in one block
	const (
		appName   = "GoApp"
		version   = "1.0.0"
		developer = "Arpit Sarang"
	)
	fmt.Println(appName, version, developer)

}
