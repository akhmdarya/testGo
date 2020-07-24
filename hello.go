package main

import "fmt"

func main() {

	// arrayVar := []int{1, 2, 3}
	// lengthVar := len(arrayVar)
	// stringVar := "Kek"

	// intVar := 1

	// float64Var := 1.0

	// fmt.Printf("%v %v %v %v  %v\n", stringVar, intVar, float64Var, arrayVar, lengthVar)

	word := "kek"

	res := PrintFunc(word)

	fmt.Println(res)
}

func PrintFunc(word string) string {

	result := fmt.Sprintf("Hello %s", word)

	return result
}
