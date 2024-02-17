package main

import (
	"fmt"
	"log"
	"strings"
)

/*

Input: s = “geeks quiz practice code”
Output: s = “code practice quiz geeks”

Input: s = “i love programming very much”
Output: s = “much very programming love i”

*/

func main() {
	fmt.Println()
	reverse()
	fmt.Println()
	log.Fatal()
	Input := "geeks quiz practice code"
	s := "i love programming very much"

	tempAr := strings.Split(Input, " ")
	tempAr2 := strings.Split(s, " ")

	newAr := []string{}

	for i := len(tempAr) - 1; i >= 0; {
		newAr = append(newAr, tempAr[i])
		i--
	}
	newAr2 := []string{}
	for j := len(tempAr2) - 1; j >= 0; {
		newAr2 = append(newAr2, tempAr2[j])
		j--
	}
	fmt.Println(strings.Join(newAr, " "))
	fmt.Println(strings.Join(newAr2, " "))
}

func reverse() {
	input := "geeks quiz practice code"

	words := strings.Fields(input) // Split the input string into words

	// Reverse the order of words in the slice
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the reversed words back into a sentence
	reversedSentence := strings.Join(words, " ")

	fmt.Println(reversedSentence)
}
