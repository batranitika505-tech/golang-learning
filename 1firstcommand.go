// Your 1st Golang Command

package main //When you run a Go program, the Go runtime looks for package main and executes its main() function.

import "fmt"

// import lets you include other packages (built-in or external).
// "fmt" stands for format, and it’s part of Go’s standard library.


func main() {
    fmt.Println("Hello")
    fmt.Println("Arpit")

	fmt.Print("Hello ")
    fmt.Println("Arpit")

	name := "Arpit"
    age := "20"
    fmt.Printf("My name is %s and I am %d years old.\n", name, age)
}

// | Specifier | Description                    | Example             |
// | --------- | ------------------------------ | ------------------- |
// | `%s`      | String                         | `"Arpit"` → Arpit   |
// | `%d`      | Integer (decimal)              | `20` → 20           |
// | `%f`      | Float                          | `3.14` → 3.140000   |
// | `%.2f`    | Float with 2 decimals          | `3.14159` → 3.14    |
// | `%t`      | Boolean                        | `true` → true       |
// | `%v`      | Default format (any value)     | `"GoLang"` → GoLang |
// | `%T`      | Print the **type** of variable | `"Arpit"` → string  |



