package main

import "fmt"

func printString(s string) {
	fmt.Println(s)
}

func appendString(s1 string, s2 string) string  {
	return s1 + s2
}

func main() {
	var w1, w2 string
	fmt.Println("Input word one: ")
	fmt.Scanln(&w1)
	fmt.Println("Input word two: ")
	fmt.Scanln(&w2)

	fmt.Println()

	fmt.Println("inputed string")
	printString(w1)
	printString(w2)

	fmt.Println()

	fmt.Println("Gabungan kedua String: ")
	sentence := appendString(w1, w2)
	printString(sentence)

	

}