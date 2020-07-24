package main

import "fmt"

func main() {

	word := "kek"

	res := PrintFunc(word)

	fmt.Println(res)
}

func PrintFunc(word string) string {

	result := fmt.Sprintf("Hello %s", word)

	return result
}
