package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	size := 0
	for i := range arg {
		size = i + 1
	}
	if size > 0 {
			argStr := arg[0]
	for _, i := range arg[1:] {
		argStr += " " + i
	}	
	vowStr := ""
	for _, v := range argStr {
		if isVowel(v) {
			vowStr += string(v) 	
		}
	}
	for i := range vowStr {
		size = i+1
	}
	finalStr := ""
	for _, v := range argStr  {
		if isVowel(v) {
			v = rune(vowStr[size-1])
			size--
		}
		finalStr += string(v)
	}
	for _, v := range finalStr {
		z01.PrintRune(v)
	}
	} 
	z01.PrintRune('\n')
}
func isVowel(v rune) bool {	
		return v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U'
}