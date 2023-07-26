package main

import "fmt"

func printString(s string) {
	fmt.Println(s)
}

func appendString(s1 string, s2 string) string  {
	return s1 + s2
}

func printNTimes(s string, n int)  {
	for i := 0 ; i < n ; i++ {
		fmt.Println(s)
	} 
	
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

	var n int 
	fmt.Println()
	fmt.Println("Input how many will print combined string: ")
	fmt.Scanf("%d", &n)
	printNTimes(sentence, n)

}