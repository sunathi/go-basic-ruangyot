package calculator

import "fmt"

// name_func({admit variable} {admit variable type}) {return variable type}
// Add => public / add => private (in packege only)
func Add(a, b int) int {
	fmt.Println("multiply", multiply(a, b))
	return a + b
}