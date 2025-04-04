// Description: This program converts the boiling point of water from Kelvin to Celsius.
package main

import "fmt"

// Declaring a constant variable for the boiling point of water in Kelvin
const boilingK float64 = 373.15

func main() {
	// I can use Gopher's friend, the walrus operator (:=) to declare and assign a value to a variable
	// The walrus operator is only available inside code blocks
	C := boilingK - 273.15

	// Printing the result to the console; using %T to print the type of the variable
	fmt.Printf("The boiling point of water is = %.2f°C (%T)\n", C, C)
}
